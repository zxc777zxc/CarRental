package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"

	"CarRental/rental-service/config"
	"CarRental/rental-service/infrastructure/repository"
	grp "CarRental/rental-service/internal/delivery/grpc"
	"CarRental/rental-service/internal/usecase"
	pb "CarRental/rental-service/proto"

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

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	repo := repository.NewRentalRepo(db)
	uc := usecase.NewRentalUsecase(repo)
	handler := grp.NewRentalHandler(uc)

	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRentalServiceServer(grpcServer, handler)
	log.Println("Rental Service running on :50056")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
