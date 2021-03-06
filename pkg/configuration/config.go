package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration
type Configuration struct {
	ListenPort      string `default:":9670" split_words:"true"`
	RootURL         string `default:"/go-auth" split_words:"true"`
	OriginHost      string `default:"delivery-platform-backend" split_words:"true"`
	Timeout         int64  `default:"60000" split_words:"true"`
	Addr            string `default:"localhost" split_words:"true"`
	MariaDBAddr     string `default:"localhost" split_words:"true"`
	MariaDBPort     string `default:"3307" split_words:"true"`
	MariaDBUser     string `default:"goauth" split_words:"true"`
	MariaDBPassword string `default:"THTqAOELuFckJZZaBP7Z" split_words:"true"`
	MariaDBDatabase string `default:"dbauth" split_words:"true"`
	LimitQuery      int64  `default:"10" split_words:"true"`
	ClientSecret    string `default:"D3l1v3ryPlatformSecretdev" split_words:"true"`
	TokenLifeTime   int64  `default:"10800" split_words:"true"`
	RedisDBAddr     string `default:"54.179.180.182" split_words:"true"`
	RedisPort       string `default:"6379" split_words:"true"`
	RedisDBPassword string `default:"" split_words:"true"`
	ServiceEmail    string `default:"http://54.179.180.182:9671/notif-email/send" split_words:"true"`
}

// Config .
var Config Configuration

// LoadConfig .
func LoadConfig() {
	if err := envconfig.Process("DP", &Config); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
