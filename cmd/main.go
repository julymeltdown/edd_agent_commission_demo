package main

import (
	"context"
	"edd_agent_commission/internal/config"
	"edd_agent_commission/internal/consumer"
	"edd_agent_commission/internal/handler"
	"edd_agent_commission/internal/repository"
	"edd_agent_commission/internal/service"
	"log"
	"strings"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// PostgreSQL 연결
	pool, err := pgxpool.Connect(context.Background(), cfg.PostgresqlDSN)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer pool.Close()

	// 레포지토리, 서비스, 핸들러 초기화
	commissionRepository := repository.NewCommissionRepository(pool)
	commissionService := service.NewCommissionService(commissionRepository)
	commissionHandler := handler.NewCommissionHandler(commissionService)

	// Echo 서버 시작
	e := echo.New()
	e.GET("/commissions/:agentName", commissionHandler.GetCommissionsByAgentName)
	// ... (다른 핸들러 추가)

	go consumer.StartConsumer(context.Background(), strings.Split(cfg.KafkaBrokers, ","), cfg.KafkaTopic, commissionService)

	// HTTP 서버 시작
	log.Fatal(e.Start(":" + cfg.HTTPServerPort))
}
