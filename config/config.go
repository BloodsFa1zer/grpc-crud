package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	UserDBPassword string `env:"UserDatabasePassword"`
	UserDBName     string `env:"UserDatabaseName"`
	DBName         string `env:"DatabaseName"`
	DriverDBName   string `env:"DriverDatabaseName"`
	Addr           string `env:"GrpcAdrr"`
}

func LoadENV(filename string) *Config {
	err := godotenv.Load(filename)
	if err != nil {
		log.Panic().Err(err).Msg(" does not load .env")
	}
	log.Info().Msg("successfully load .env")
	cfg := Config{}
	return &cfg

}

func (cfg *Config) ParseENV() {

	err := env.Parse(cfg)
	if err != nil {
		log.Panic().Err(err).Msg(" unable to parse environment variables")
	}
	log.Info().Msg("successfully parsed .env")
}

//type JwtCustomClaims struct {
//	ID   int64  `json:"ID"`
//	Name string `json:"name"`
//	Role string `json:"role"`
//	jwt.RegisteredClaims
//}

//func NewConfig() echojwt.Config {
//	cfg := LoadENV("config/.env")
//	cfg.ParseENV()
//
//	Config := echojwt.Config{
//		NewClaimsFunc: func(c echo.Context) jwt.Claims {
//			return new(JwtCustomClaims)
//		},
//		SigningKey: []byte(cfg.SigningKey),
//	}
//	return Config
//}
