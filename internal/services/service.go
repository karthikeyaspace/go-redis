package services

import (
	"github.com/redis/go-redis/v9"
)

type Service interface {
	UpdateScore(userId string, score int) error
	GetLeaderboard() ([]map[string]interface{}, error)
}

type service struct {
	client *redis.Client
}

func NewService(client *redis.Client) Service {
	return &service{client: client}
}

func (s *service) UpdateScore(userId string, score int) error {
	return nil
}

func (s *service) GetLeaderboard() ([]map[string]interface{}, error) {
	return nil, nil
}
