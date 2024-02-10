package net

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"url-shortener/internal/config"
	"url-shortener/internal/handlers"
	"url-shortener/internal/net/auth"
	"url-shortener/internal/storage"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	cfg      *config.Config
	log      *slog.Logger
	mongo    *storage.Mongo
	memcache *storage.Memcached
	app      *fiber.App
}

func Init() *App {
	cfg := config.NewConfig()
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	auth.SetSecret(cfg.JwtSecret)
	mongo := storage.NewMongo(cfg)
	if _, err := mongo.Ping(); err != nil {
		log.Fatal(err)	
	}
	return &App{
		cfg:      cfg,
		log:      slog.Default(),
		mongo:    mongo,
		memcache: storage.NewMemcached(cfg),
		app:      app,
	}
}

func (a *App) Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	go func() {
		<-quit
		a.log.Info("Server shutting down...")
		if err := a.app.Shutdown(); err != nil {
			a.log.Error("Server shutdown error: " + err.Error())
		}
		if _, err := a.mongo.Disconnect(); err != nil {
			a.log.Error("Database shutdown error: " + err.Error())
		}
		if err := a.memcache.Disconnect(); err != nil {
			a.log.Error("Cache shutdown error: " + err.Error())
		}
	}()

	a.ConfigureRoutes()
	if err := a.app.Listen(fmt.Sprintf("%s:%s", a.cfg.Host, a.cfg.Port)); err != nil {
		a.log.Error("Server error: " + err.Error())
	}
}

func (a *App) ConfigureRoutes() {
	a.app.Get("/:alias", func(c *fiber.Ctx) error {
		return handlers.HandleRedirect(c, a.memcache, a.mongo)
	})
	a.app.Post("/", func(c *fiber.Ctx) error {
		return handlers.HandleSignUp(c, a.mongo)
	})
	a.app.Patch("/", func(c *fiber.Ctx) error {
		return handlers.HandleSignIn(c, a.mongo)
	})
	a.app.Get("/", func(c *fiber.Ctx) error {
		return handlers.HandleShorten(c, a.memcache, a.mongo)
	})
	a.app.Delete("/", func(c *fiber.Ctx) error {
		return handlers.HandleDelete(c, a.memcache, a.mongo)
	})
}
