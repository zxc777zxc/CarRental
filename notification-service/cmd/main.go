package main

import (
	"CarRental/notification-service/config"
	"CarRental/notification-service/infrastructure/email"
	nat "CarRental/notification-service/infrastructure/nats"
	gpc "CarRental/notification-service/internal/delivery/grpc"
	pb "CarRental/notification-service/proto"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net"
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
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("failed to start metrics server: %v", err)
		}
	}()

	cfg := config.Load()

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatal("failed to connect to NATS:", err)
	}
	defer nc.Close()

	sender := email.NewEmailSender(cfg)

	nat.Subscribe(nc, sender)

	lis, err := net.Listen("tcp", ":50058")
	if err != nil {
		log.Fatalf("failed to listen on port 50058: %v", err)
	}

	notificationHandler := gpc.NewNotificationHandler(sender)
	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, notificationHandler)

	log.Println("Notification Service listening on gRPC :50058 and for NATS email events...")

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	select {}
}
