package config

import "github.com/spf13/viper"

type PortManager struct {
	RunnerPort string `mapstructure:"PORTNO"`
}

type MongoDataBase struct {
	MongoDbURL    string `mapstructure:"MONGODB_URL"`
	DataBase      string `mapstructure:"MONGODB_DATABASE"`
}

type KafkaConfigs struct {
	KafkaPort              string `mapstructure:"KAFKA_PORT"`
	KafkaTopicNotification string `mapstructure:"KAFKA_TOPIC_1"`
}

type Config struct {
	PortMngr PortManager
	MongoDB  MongoDataBase
	Kafka    KafkaConfigs
}

func LoadConfig() (*Config, error) {
	var portmngr PortManager
	var MongoDb MongoDataBase
	var kafka KafkaConfigs

	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&portmngr)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&MongoDb)
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&kafka)
	if err != nil {
		return nil, err
	}

	config := Config{PortMngr: portmngr, MongoDB: MongoDb, Kafka: kafka}
	return &config, nil

}
