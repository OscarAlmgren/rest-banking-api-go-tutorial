package domain

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SaleRepositoryDb struct {
	client *mongo.Client
}

func NewSaleRepositoryDb() SaleRepositoryDb {
	clientOptions := options.Client().ApplyURI("mongodb://root:oscar-camp-tutorial@192.168.205.6:27017/?authSource=admin&readPreference=primary&ssl=false&directConnection=true")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("MongoDB connect error.")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("MongoDB ping error.")
	}

	return SaleRepositoryDb{client: client}
}

// Interface implementations from here
func (s SaleRepositoryDb) FindOne() (*Sale, error) {
	var sale Sale

	collection := s.client.Database("banking-api").Collection("sales")
	filter := bson.D{
		{Key: "item", Value: "abc"},
	}
	err := collection.FindOne(context.TODO(), filter).Decode(&sale)
	if err != nil {
		log.Error().Str("Error", err.Error()).Msg("FindOne error in MongoDB.")
		return nil, err
	}
	log.Info().Str("action", "findOne").Msg("FindOne success")
	return &sale, nil
}

func (s SaleRepositoryDb) FindAll() ([]Sale, error) {
	var sales []Sale

	collection := s.client.Database("banking-api").Collection("sales")
	filter := bson.D{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error().Str("Error", err.Error()).Msg("FindAll error in MongoDB.")
		return nil, err
	}
	if err = cursor.All(context.TODO(), &sales); err != nil { // loop cursor -> map each to one "sale"
		log.Panic().Str("Error", err.Error()).Msg("Panic from FindAll cursor check")
	}

	return sales, nil
}

func (s SaleRepositoryDb) Create(sale Sale) (string, error) {
	collection := s.client.Database("banking-api").Collection("sales")

	inserted, err := collection.InsertOne(context.TODO(), sale)
	if err != nil {
		log.Error().Str("Error", err.Error()).Msg("Create error in Create(sale Sale).")
		return "Create error in Create(sale Sale).", err
	}
	resp := inserted.InsertedID.(primitive.ObjectID).Hex()

	return resp, nil
}

func (s SaleRepositoryDb) Delete(id string) (int64, error) {
	collection := s.client.Database("banking-api").Collection("sales")

	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal().Str("Error", err.Error()).Msg("primitive.ObjectIDFromHex(id) from Delete(id string) error")
		return -1, err
	}
	filter := bson.M{"_id": idPrimitive}
	deleted, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Panic().Str("Error", err.Error()).Msg("Unable to delete from Delete(id string)")
		return -1, err
	}
	return deleted.DeletedCount, nil
}
