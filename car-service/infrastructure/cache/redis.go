package cache

import (
	"CarRental/car-service/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type CarCache struct {
	client *redis.Client
}

func NewCarCache(addr string) *CarCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &CarCache{client: rdb}
}

func (c *CarCache) SetCar(ctx context.Context, car *domain.Car) error {
	data, _ := json.Marshal(car)
	return c.client.Set(ctx, fmt.Sprintf("car:%d", car.ID), data, 12*time.Hour).Err()
}

func (c *CarCache) GetCar(ctx context.Context, id int64) (*domain.Car, error) {
	val, err := c.client.Get(ctx, fmt.Sprintf("car:%d", id)).Result()
	if err != nil {
		return nil, err
	}
	var car domain.Car
	err = json.Unmarshal([]byte(val), &car)
	return &car, err
}

func (c *CarCache) SetCarList(ctx context.Context, cars []*domain.Car) error {
	data, _ := json.Marshal(cars)
	return c.client.Set(ctx, "car:list", data, 12*time.Hour).Err()
}

func (c *CarCache) GetCarList(ctx context.Context) ([]*domain.Car, error) {
	val, err := c.client.Get(ctx, "car:list").Result()
	if err != nil {
		return nil, err
	}
	var cars []*domain.Car
	err = json.Unmarshal([]byte(val), &cars)
	return cars, err
}
