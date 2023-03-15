package app

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/config"
	"forum/internal/controller/http/handlers"
	"forum/internal/infrastructure/repository"
	"forum/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	cfg *config.Config

	db *sql.DB

	httpServer  *http.Server
	httpHandler *handlers.Controller
}

func InitApp(cfg *config.Config) (*app, error) {
	db, err := repository.RunDb()
	if err != nil {
		return nil, err
	}

	repos := repository.NewRepository(db)

	service := service.NewService(repos)

	control := handlers.NewContoller(service)

	return &app{
		cfg:         cfg,
		db:          db,
		httpHandler: control,
	}, nil
}

func (a *app) Run() {
	go func() {
		if err := a.startHttpServer(); err != nil {
			log.Fatal(err)
			return
		}
	}()

	fmt.Printf("http server started on http://localhost:%s/\n", a.cfg.Http.Port)
	// log.Println("http server started on", a.cfg.Http.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan
	fmt.Println()
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
		return
	}

	if err := a.db.Close(); err != nil {
		log.Println(err)
	} else {
		log.Println("db closed")
	}
}

func (a *app) startHttpServer() error {
	router := a.httpHandler.InitRouter()

	a.httpServer = &http.Server{
		Addr:           ":" + a.cfg.Http.Port,
		Handler:        router,
		ReadTimeout:    time.Second * time.Duration(a.cfg.Http.ReadTimeout),
		WriteTimeout:   time.Second * time.Duration(a.cfg.Http.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}

	return a.httpServer.ListenAndServe()
}
