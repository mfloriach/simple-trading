package trade

import (
	"context"
	"fmt"

	"example.com/grpc-todo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "trades"

type cacheRepository struct {
	collection *mongo.Collection
}

func NewTradeCacheRepository() TradeRepository {
	return &cacheRepository{collection: utils.GetMongoClient().Collection(collectionName)}
}

func NewTradeAdminCacheRepository() TradeRepositoryAdmin {
	return &cacheRepository{collection: utils.GetMongoClient().Collection(collectionName)}
}

func (s *cacheRepository) Set(ctx context.Context, t Trade) error {
	fmt.Println(t)
	_, err := s.collection.InsertOne(ctx, t)

	return err
}

func (s *cacheRepository) GetByUserID(ctx context.Context, userID int, action Action) ([]Trade, error) {
	var results []Trade
	return results, nil
}

func (s *cacheRepository) Get(ctx context.Context, symbol string, action Action) ([]Trade, error) {
	c, err := s.collection.Find(ctx, bson.D{{"symbol", symbol}, {"action", action}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	var traders []Trade
	if err = c.All(ctx, &traders); err != nil {
		return nil, err
	}

	return traders, nil
}

func (s *cacheRepository) Delete(ctx context.Context, id string) error {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": idPrimitive})

	return err
}
