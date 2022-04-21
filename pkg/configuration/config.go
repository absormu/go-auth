package configuration

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Configuration
type Configuration struct {
	ListenPort      string `default:":9670" split_words:"true"`
	RootURL         string `default:"/delivery-platform-backend" split_words:"true"`
	OriginHost      string `default:"delivery-platform-backend" split_words:"true"`
	Timeout         int64  `default:"60000" split_words:"true"`
	Addr            string `default:"localhost" split_words:"true"`
	MariaDBAddr     string `default:"localhost" split_words:"true"`
	MariaDBPort     string `default:"3306" split_words:"true"`
	MariaDBUser     string `default:"goauth" split_words:"true"`
	MariaDBPassword string `default:"THTqAOELuFckJZZaBP7Z" split_words:"true"`
	MariaDBDatabase string `default:"dbauth" split_words:"true"`
	LimitQuery      int64  `default:"10" split_words:"true"`
	ClientSecret    string `default:"D3l1v3ryPlatformSecretdev" split_words:"true"`
	TokenLifeTime   int64  `default:"10800" split_words:"true"`
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
