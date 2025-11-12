package store

import (
	db "go-boilerplate/store/sqlc"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	PostgresDB  db.DBInterface
	RedisClient *redis.Client
	Querier     *db.Queries
}

func NewStoreBuilder() *StoreBuilder {
	return &StoreBuilder{}
}

type StoreBuilder struct {
	postgresDB  db.DBInterface
	redisClient *redis.Client
	querier     *db.Queries
}

func (b *StoreBuilder) WithPostgres(postgresDB db.DBInterface) *StoreBuilder {
	b.postgresDB = postgresDB
	return b
}

func (b *StoreBuilder) WithRedis(client *redis.Client) *StoreBuilder {
	b.redisClient = client
	return b
}

func (b *StoreBuilder) WithQuerier(querier *db.Queries) *StoreBuilder {
	b.querier = querier
	return b
}

func (b *StoreBuilder) Build() *Store {
	return &Store{
		PostgresDB:  b.postgresDB,
		RedisClient: b.redisClient,
		Querier:     b.querier,
	}
}
