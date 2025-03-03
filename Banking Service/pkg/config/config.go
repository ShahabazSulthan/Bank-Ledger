package config

import "github.com/spf13/viper"

type PortManager struct {
	PortNo         string `mapstructure:"PORTNO"`
	PostNrelSvcUrl string `mapstructure:"POSTNREL_SVC_URL"`
}

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBHost     string `mapstructure:"DBHOST"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Config struct {
	PortMngr PortManager
	DB       DataBase
	Kafka    KafkaConfigs
}

type KafkaConfigs struct {
	KafkaPort              string `mapstructure:"KAFKA_PORT"`
	KafkaTopicNotification string `mapstructure:"KAFKA_TOPIC_2"`
}

func LoadConfig() (*Config, error) {
	var portmanager PortManager
	var db DataBase
	var kafka KafkaConfigs

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&portmanager)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&db)
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&kafka)
	if err != nil {
		return nil, err
	}

	config := Config{PortMngr: portmanager, DB: db, Kafka: kafka}
	return &config, nil
}
