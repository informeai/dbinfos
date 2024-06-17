package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/informeai/dbinfos/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type dbInfoMongoAdapter struct {
	uri    string
	client *mongo.Client
	config config.DBInfoMongoConfigure
}

// NewDBInfoMongoDB return instance the mongo adapter
func NewDBInfoMongoDB(config config.DBInfoMongoConfigure) *dbInfoMongoAdapter {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?ssl=false&authSource=admin", config.Username, config.Password, config.Host, config.Port, config.Database)
	return &dbInfoMongoAdapter{uri: uri, config: config}
}

func (m *dbInfoMongoAdapter) marshaler(infos any) (string, error) {
	infBytes, err := json.Marshal(infos)
	if err != nil {
		return "", err
	}
	return string(infBytes), nil
}

// Connect execute connection in database mongo
func (m *dbInfoMongoAdapter) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.uri))
	if err != nil {
		return err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	m.client = client
	return nil
}

// Save execute persistence the dbinfos in database mongo
func (m *dbInfoMongoAdapter) Save(topic string, infos any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	now := time.Now()
	collection := m.client.Database(m.config.Database).Collection(m.config.Collection)
	_, err := collection.InsertOne(ctx, bson.M{"topic": topic, "info": infos, "created_at": now, "updated_at": now})
	if err != nil {
		return err
	}
	return nil
}

// Disconnect execute disconnect in database mongo
func (m *dbInfoMongoAdapter) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := m.client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
