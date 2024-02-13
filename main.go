package main

//go:generate go run -mod=mod ./ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alexferrari88/gohn/pkg/gohn"
	"github.com/breadchris/hn/ent"
	"github.com/breadchris/hn/ent/item"
	"github.com/breadchris/hn/ent/migrate"
	"github.com/lmittmann/tint"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	logFn := func(params ...any) {
		slog.Debug("ent", params)
	}
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
	client, err := ent.Open(
		dialect.SQLite,
		"file:hn.db?cache=shared&_fk=1",
		ent.Log(logFn),
		ent.Debug(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	hn, err := gohn.NewClient(nil)
	if err != nil {
		slog.Error("failed creating hn client", "err", err)
		return
	}

	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx,
		// TODO breadchris do not configure this for production
		// use https://entgo.io/docs/versioned-migrations
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := &cli.App{
		Name:  "hn",
		Usage: "A Hacker News client and server",
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Start the server",
				Action: func(c *cli.Context) error {
					srv := handler.NewDefaultServer(NewSchema(client))
					http.Handle("/",
						playground.Handler("Item", "/query"),
					)
					http.Handle("/query", srv)
					log.Println("listening on :8081")
					if err := http.ListenAndServe(":8081", nil); err != nil {
						log.Fatal("http server terminated", err)
					}
					return nil
				},
			},
			{
				Name:    "replicate",
				Aliases: []string{"r"},
				Usage:   "Replicate data",
				Action: func(c *cli.Context) error {
					const numWorkers = 10
					items := make(chan int, 100)
					var wg sync.WaitGroup
					errChan := make(chan error, 1)
					itemChan := make(chan *gohn.Item, 100)

					for i := 0; i < numWorkers; i++ {
						wg.Add(1)
						go worker(ctx, hn, client, items, &wg, itemChan, errChan)
					}

					go func() {
						for i := 39346345; i > 0; i-- {
							select {
							case items <- i:
							case err := <-errChan:
								slog.Info("error processing item", "err", err)
							}
						}
						close(items)
					}()

					go func() {
						wg.Add(1)
						defer wg.Done()
						for i := range itemChan {
							if err := upsertItem(ctx, client, i); err != nil {
								errChan <- err
								return
							}
						}
					}()

					wg.Wait()

					select {
					case err := <-errChan:
						log.Fatalf("error processing item: %v", err)
					default:
					}
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func worker(ctx context.Context, hn *gohn.Client, client *ent.Client, items <-chan int, wg *sync.WaitGroup, itemsChan chan<- *gohn.Item, errChan chan<- error) {
	defer wg.Done()
	for itemID := range items {
		i, err := hn.Items.Get(ctx, itemID)
		if err != nil {
			errChan <- err
			return
		}
		itemsChan <- i
	}
}

func upsertItem(ctx context.Context, client *ent.Client, i *gohn.Item) error {
	if i == nil {
		slog.Warn("item is nil")
		return nil
	}

	strID := fmt.Sprintf("%d", *i.ID)
	existingItem, err := client.Item.Query().Where(item.IDEQ(strID)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	if existingItem == nil {
		create := client.Item.Create()
		if i.ID != nil {
			create.SetID(strID)
		}
		if i.Deleted != nil {
			create.SetDeleted(*i.Deleted)
		} else {
			create.SetDeleted(false)
		}
		if i.Type != nil {
			create.SetType(item.Type(*i.Type))
		}
		if i.By != nil {
			create.SetBy(*i.By)
		}
		if i.Time != nil {
			create.SetTime(*i.Time)
		} else {
			create.SetTime(0)
		}
		if i.Text != nil {
			create.SetText(*i.Text)
		}
		if i.Dead != nil {
			create.SetDead(*i.Dead)
		} else {
			create.SetDead(false)
		}
		if i.Parent != nil {
			create.SetParent(*i.Parent)
		}
		if i.Poll != nil {
			create.SetPoll(*i.Poll)
		}
		if i.Kids != nil {
			create.SetKids(*i.Kids)
		}
		if i.URL != nil {
			create.SetURL(*i.URL)
		}
		if i.Score != nil {
			create.SetScore(*i.Score)
		}
		if i.Title != nil {
			create.SetTitle(*i.Title)
		}
		if i.Parts != nil {
			create.SetParts(*i.Parts)
		}
		if i.Descendants != nil {
			create.SetDescendants(*i.Descendants)
		}
		_, err := create.Save(ctx)
		if err != nil {
			return err
		}
	} else {
		update := client.Item.UpdateOneID(strID)
		if i.Deleted != nil {
			update.SetDeleted(*i.Deleted)
		} else {
			update.SetDeleted(false)
		}
		if i.Type != nil {
			update.SetType(item.Type(*i.Type))
		}
		if i.By != nil {
			update.SetBy(*i.By)
		}
		if i.Time != nil {
			update.SetTime(*i.Time)
		} else {
			update.SetTime(0)
		}
		if i.Text != nil {
			update.SetText(*i.Text)
		}
		if i.Dead != nil {
			update.SetDead(*i.Dead)
		} else {
			update.SetDead(false)
		}
		if i.Parent != nil {
			update.SetParent(*i.Parent)
		}
		if i.Poll != nil {
			update.SetPoll(*i.Poll)
		}
		if i.Kids != nil {
			update.SetKids(*i.Kids)
		}
		if i.URL != nil {
			update.SetURL(*i.URL)
		}
		if i.Score != nil {
			update.SetScore(*i.Score)
		}
		if i.Title != nil {
			update.SetTitle(*i.Title)
		}
		if i.Parts != nil {
			update.SetParts(*i.Parts)
		}
		if i.Descendants != nil {
			update.SetDescendants(*i.Descendants)
		}
		_, err := update.Save(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
