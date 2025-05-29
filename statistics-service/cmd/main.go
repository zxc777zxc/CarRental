package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"CarRental/statistics-service/config"
	nat "CarRental/statistics-service/infrastructure/nats"
	"CarRental/statistics-service/infrastructure/repository"
	grp "CarRental/statistics-service/internal/delivery/grpc"
	"CarRental/statistics-service/internal/usecase"
	pb "CarRental/statistics-service/proto"
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

	db, _ := sql.Open("postgres", cfg.DBUrl)
	defer db.Close()

	repo := repository.NewStatisticsRepo(db)
	uc := usecase.NewStatisticsUsecase(repo)

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	nat.Subscribe(nc, uc)

	lis, _ := net.Listen("tcp", ":50055")
	s := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(s, grp.NewHandler(uc))

	log.Println("Statistics Service running at :50055")
	s.Serve(lis)
}
