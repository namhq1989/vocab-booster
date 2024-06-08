package config

import "errors"

type (
	Server struct {
		RestPort string
		GRPCPort string

		AppName      string
		Environment  string
		IsEnvRelease bool
		Debug        bool

		// Single Sign On
		SSOGoogleClientID     string
		SSOGoogleClientSecret string

		// Authentication
		AccessTokenSecret  string
		RefreshTokenSecret string
		AccessTokenTTL     int // seconds
		RefreshTokenTTL    int // seconds

		// MongoDB
		MongoURL    string
		MongoDBName string

		// Meilisearch
		MeilisearchHost   string
		MeilisearchAPIKey string

		// Redis
		CachingRedisURL string

		// Queue
		QueueRedisURL    string
		QueueUsername    string
		QueuePassword    string
		QueueConcurrency int

		// Sentry
		SentryDSN     string
		SentryMachine string

		// OpenAI
		OpenAIAPIKey string

		// Ably
		AblyAPIKey string
	}
)

func Init() Server {
	cfg := Server{
		RestPort: ":3000",
		GRPCPort: ":3001",

		AppName:     getEnvStr("APP_NAME"),
		Environment: getEnvStr("ENVIRONMENT"),
		Debug:       getEnvBool("DEBUG"),

		SSOGoogleClientID:     getEnvStr("SSO_GOOGLE_CLIENT_ID"),
		SSOGoogleClientSecret: getEnvStr("SSO_GOOGLE_CLIENT_SECRET"),

		AccessTokenSecret:  getEnvStr("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: getEnvStr("REFRESH_TOKEN_SECRET"),
		AccessTokenTTL:     getEnvInt("ACCESS_TOKEN_TTL"),
		RefreshTokenTTL:    getEnvInt("REFRESH_TOKEN_TTL"),

		MongoURL:    getEnvStr("MONGO_URL"),
		MongoDBName: getEnvStr("MONGO_DB_NAME"),

		MeilisearchHost:   getEnvStr("MEILISEARCH_HOST"),
		MeilisearchAPIKey: getEnvStr("MEILISEARCH_API_KEY"),

		CachingRedisURL: getEnvStr("CACHING_REDIS_URL"),

		QueueRedisURL:    getEnvStr("QUEUE_REDIS_URL"),
		QueueUsername:    getEnvStr("QUEUE_USERNAME"),
		QueuePassword:    getEnvStr("QUEUE_PASSWORD"),
		QueueConcurrency: getEnvInt("QUEUE_CONCURRENCY"),

		SentryDSN:     getEnvStr("SENTRY_DSN"),
		SentryMachine: getEnvStr("SENTRY_MACHINE"),

		OpenAIAPIKey: getEnvStr("OPENAI_API_KEY"),

		AblyAPIKey: getEnvStr("ABLY_API_KEY"),
	}
	cfg.IsEnvRelease = cfg.Environment == "release"

	// validation
	if cfg.Environment == "" {
		panic(errors.New("missing ENVIRONMENT"))
	}

	if cfg.MongoURL == "" {
		panic(errors.New("missing MONGO_URL"))
	}
	if cfg.MongoDBName == "" {
		panic(errors.New("missing MONGO_DB_NAME"))
	}
	if cfg.MongoDBName == "" {
		panic(errors.New("missing MONGO_DB_NAME"))
	}

	if cfg.MeilisearchHost == "" {
		panic(errors.New("missing MEILISEARCH_HOST"))
	}
	if cfg.MeilisearchAPIKey == "" {
		panic(errors.New("missing MEILISEARCH_API_KEY"))
	}

	if cfg.CachingRedisURL == "" {
		panic(errors.New("missing CACHING_REDIS_URL"))
	}

	if cfg.QueueRedisURL == "" {
		panic(errors.New("missing QUEUE_REDIS_URL"))
	}

	if cfg.AccessTokenSecret == "" {
		panic(errors.New("missing ACCESS_TOKEN_SECRET"))
	}
	if cfg.RefreshTokenSecret == "" {
		panic(errors.New("missing REFRESH_TOKEN_SECRET"))
	}

	if cfg.OpenAIAPIKey == "" {
		panic(errors.New("missing OPENAI_API_KEY"))
	}

	if cfg.AblyAPIKey == "" {
		panic(errors.New("missing ABLY_API_KEY"))
	}

	return cfg
}
