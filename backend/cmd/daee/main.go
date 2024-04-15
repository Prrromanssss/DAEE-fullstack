package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Prrromanssss/DAEE-fullstack/internal/agent"
	"github.com/Prrromanssss/DAEE-fullstack/internal/config"
	"github.com/Prrromanssss/DAEE-fullstack/internal/http-server/handlers"
	"github.com/Prrromanssss/DAEE-fullstack/internal/http-server/middleware"
	"github.com/Prrromanssss/DAEE-fullstack/internal/orchestrator"
	"github.com/Prrromanssss/DAEE-fullstack/internal/storage"
	"github.com/Prrromanssss/DAEE-fullstack/lib/logger/handlers/slogpretty"
	"github.com/Prrromanssss/DAEE-fullstack/lib/logger/logcleaner"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't get pwd")
	}
	rootPath := filepath.Dir(filepath.Dir(path))
	logPath := fmt.Sprintf("%s/daee.log", rootPath)

	// Configuration log file
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Log file is not found in environment")
	} else {
		log.SetOutput(file)
	}
	defer file.Close()

	go logcleaner.CleanLog(10*time.Minute, logPath, 100)

	// Load env variables
	err = godotenv.Load(fmt.Sprintf("%s/.env", filepath.Dir(rootPath)))
	if err != nil {
		log.Fatalf("Can't load env variables: %v", err)
	}

	// Load config
	cfg := config.MustLoad()

	dbCfg := storage.NewStorage(cfg.StorageURL)

	agentAgregator, err := agent.NewAgentAgregator(
		cfg.RabbitMQURL,
		dbCfg,
		cfg.QueueForSendToAgents,
		cfg.QueueForConsumeFromAgents,
	)
	if err != nil {
		log.Fatalf("Agent Agregator Error: %v", err)
	}

	go agent.AgregateAgents(agentAgregator)

	// Reload computing expressions
	err = orchestrator.ReloadComputingExpressions(dbCfg, agentAgregator)
	if err != nil {
		log.Fatalf("Can't reload computin expressions: %v", err)
	}

	// Delete previous agents
	err = dbCfg.DB.DeleteAgents(context.Background())
	if err != nil {
		log.Fatalf("Can't delete previous agents: %v", err)
	}

	// Create Agent1
	agent1, err := agent.NewAgent(
		cfg.RabbitMQURL,
		dbCfg,
		cfg.QueueForSendToAgents,
		cfg.QueueForConsumeFromAgents,
		5,
		200,
	)
	if err != nil {
		log.Fatalf("Can't create agent1: %v", err)
	}

	go agent.AgentService(agent1)

	// Create Agent2
	agent2, err := agent.NewAgent(
		cfg.RabbitMQURL,
		dbCfg,
		cfg.QueueForSendToAgents,
		cfg.QueueForConsumeFromAgents,
		5,
		200,
	)
	if err != nil {
		log.Fatalf("Can't create agent2: %v", err)
	}

	go agent.AgentService(agent2)

	// Create Agent3
	agent3, err := agent.NewAgent(
		cfg.RabbitMQURL,
		dbCfg,
		cfg.QueueForSendToAgents,
		cfg.QueueForConsumeFromAgents,
		5,
		200,
	)
	if err != nil {
		log.Fatalf("Can't create agent2: %v", err)
	}

	go agent.AgentService(agent3)

	// Configuration http server
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Post("/expressions", middleware.MiddlewareAgentAgregatorAndDBConfig(
		handlers.HandlerCreateExpression,
		dbCfg,
		agentAgregator,
	))
	v1Router.Get("/expressions", middleware.MiddlewareApiConfig(handlers.HandlerGetExpressions, dbCfg))

	v1Router.Get("/operations", middleware.MiddlewareApiConfig(handlers.HandlerGetOperations, dbCfg))
	v1Router.Patch("/operations", middleware.MiddlewareApiConfig(handlers.HandlerUpdateOperation, dbCfg))

	v1Router.Get("/agents", middleware.MiddlewareApiConfig(handlers.HandlerGetAgents, dbCfg))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.Address,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	log.Printf("Server starting on port %v", 3000)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
