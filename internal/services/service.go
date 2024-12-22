package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/karthikeyaspace/game-leaderboard/internal/utils"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	CreatePlayerService(name string) (string, error)
	UpdateScoreService(userId string, score int) error
	GetLeaderboardService(limit string) ([]map[string]interface{}, error)
}

type service struct {
	client *redis.Client
}

var ctx = context.Background()

func NewService(client *redis.Client) Service {
	return &service{client: client}
}

func (s *service) CreatePlayerService(name string) (string, error) {
	userId := utils.GenerateID(5)
	key := fmt.Sprintf("user:%s", userId)
	_, err := s.client.HMSet(ctx, key, "name", name, "score", 0).Result()
	if err != nil {
		return "", err
	}
	return userId, nil
}

func (s *service) UpdateScoreService(userId string, score int) error {
	key := fmt.Sprintf("user:%s", userId)

	if _, err := s.client.HSet(ctx, key, "score", score).Result(); err != nil {
		return err
	}

	_, err := s.client.ZAdd(ctx, "leaderboard", redis.Z{
		Score:  float64(score),
		Member: userId,
	}).Result()

	if err != nil {
		return err
	}

	return nil

}

func (s *service) GetLeaderboardService(limit string) ([]map[string]interface{}, error) {
	lim, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}

	players, err := s.client.ZRevRangeWithScores(ctx, "leaderboard", 0, int64(lim-1)).Result()

	if err != nil {
		return nil, err
	}

	var leaderboard []map[string]interface{}

	for _, player := range players {
		userID := player.Member.(string)
		userData, err := s.client.HGetAll(ctx, fmt.Sprintf("user:%s", userID)).Result()
		if err != nil {
			continue
		}

		leaderboard = append(leaderboard, map[string]interface{}{
			"userId": userID,
			"name":   userData["name"],
			"score":  int(player.Score),
		})
	}

	return leaderboard, nil
}
