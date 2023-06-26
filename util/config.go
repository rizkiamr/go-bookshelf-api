package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	ServiceName                    string `mapstructure:"SERVICE_NAME"`
	ServiceVersion                 string `mapstructure:"SERVICE_VERSION"`
	ServiceBasicAuthId             string `mapstructure:"SERVICE_BASIC_AUTH_ID"`
	ServiceBasicAuthSecret         string `mapstructure:"SERVICE_BASIC_AUTH_SECRET"`
	ServicePort                    string `mapstructure:"SERVICE_PORT"`
	ServiceAddress                 string `mapstructure:"SERVICE_ADDRESS"`
	ServiceBasePath                string `mapstructure:"SERVICE_BASE_PATH"`
	ServiceRateLimitQuotaPerMinute string `mapstructure:"SERVICE_RATE_LIMIT_QUOTA_PER_MINUTE"`
	DbDriver                       string `mapstructure:"DB_DRIVER"`
	DbHost                         string `mapstructure:"DB_HOST"`
	DbPort                         string `mapstructure:"DB_PORT"`
	DbName                         string `mapstructure:"DB_NAME"`
	DbUser                         string `mapstructure:"DB_USER"`
	DbPassword                     string `mapstructure:"DB_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
