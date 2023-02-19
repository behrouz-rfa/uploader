package config

import (
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"

	"github.com/stackus/dotenv"

	"uploader/internal/rpc"
	"uploader/internal/web"
)

type (
	PGConfig struct {
		Conn string `required:"true"`
	}
	MinioConfig struct {
		PublicEndPoint string `required:"true" json:"public_end_point"`
		EndPoint       string `required:"true" json:"end_point"`
		AccessKey      string `required:"true" json:"access_key"`
		SecretKey      string `required:"true" json:"secret_key"`
		Secure         bool   `required:"true" json:"secure,omitempty"`
	}

	AppConfig struct {
		Environment     string
		LogLevel        string `envconfig:"LOG_LEVEL" default:"DEBUG"`
		PG              PGConfig
		Minio           MinioConfig
		Rpc             rpc.RpcConfig
		Web             web.WebConfig
		ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"30s"`
	}
)

func InitConfig() (cfg AppConfig, err error) {
	if err = dotenv.Load(dotenv.EnvironmentFiles(os.Getenv("ENVIRONMENT"))); err != nil {
		return
	}

	err = envconfig.Process("", &cfg)

	return
}
