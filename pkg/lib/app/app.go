package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type App interface {
	Start(ctx context.Context) error
	Stop(context.Context) error
}

func Handle(name string, app App) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		fmt.Printf("[%s] Starting active the app", name)
		if err := app.Start(ctx); err != nil {
			fmt.Printf("\n[%s] failed to active the app: %v", name, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := app.Stop(ctx); err != nil {
		return fmt.Errorf("[%s] failed to stop the app: %v", name, err)
	}
	return nil
}
