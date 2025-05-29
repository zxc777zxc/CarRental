package main

import (
	"CarRental/car-service/config"
	"CarRental/car-service/infrastructure/cache"
	"CarRental/car-service/infrastructure/repository"
	grp "CarRental/car-service/internal/delivery/grpc"
	"CarRental/car-service/internal/usecase"
	pb "CarRental/car-service/proto"
	"database/sql"
	"google.golang.org/grpc"
	"log"
	"net"

	_ "github.com/lib/pq"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func main() {

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	rdb := cache.NewCarCache(cfg.RedisAddr)

	repo := repository.NewCarRepository(db)
	uc := usecase.NewCarUsecase(repo, rdb)
	handler := grp.NewCarHandler(uc)

	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCarServiceServer(server, handler)

	log.Println("CarService listening on", cfg.GRPCPort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
