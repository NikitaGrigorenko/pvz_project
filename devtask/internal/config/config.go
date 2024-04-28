package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AuthInfo struct {
	Password string `json:"AUTH_PASSWORD"`
	Username string `json:"AUTH_USERNAME"`
}

type DbInfo struct {
	Host     string `json:"DB_HOST"`
	Name     string `json:"DB_NAME"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	User     string `json:"DB_USER"`
}

type KafkaInfo struct {
	Topic string `json:"KAFKA_TOPIC"`
	Group string `json:"KAFKA_GROUP"`
}

type BrokersInfo struct {
	Brokers string `json:"BROKERS"`
}

type RedisInfo struct {
	Addr     string `json:"ADDR"`
	Password string `json:"PASSWORD"`
	DB       int    `json:"DB"`
}

type AddrInfo struct {
	Addr       string `json:"GRPC_ADDR"`
	AddrJaeger string `json:"JAEGER_ADDR"`
}

type Config struct {
	AuthInfo
	DbInfo
	Brokers []string
	KafkaInfo
	RedisInfo
	AddrInfo
}

func Read(fileName string) (Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Config{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var auth AuthInfo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&auth)
	if err != nil {
		fmt.Println("Error parsing AuthInfo:", err)
		return Config{}, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return Config{}, err
	}

	var db DbInfo
	err = decoder.Decode(&db)
	if err != nil {
		fmt.Println("Error parsing DBInfo:", err)
		return Config{}, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return Config{}, err
	}

	var brokers BrokersInfo
	err = decoder.Decode(&brokers)
	if err != nil {
		fmt.Println("Error parsing BrokersInfo:", err)
		return Config{}, err
	}
	var brokersSlice []string
	brokersS := strings.Split(brokers.Brokers, ";")
	brokersSlice = append(brokersSlice, brokersS...)

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return Config{}, err
	}

	var infoKafka KafkaInfo
	err = decoder.Decode(&infoKafka)
	if err != nil {
		fmt.Println("Error parsing KafkaInfo:", err)
		return Config{}, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return Config{}, err
	}

	var infoRedis RedisInfo
	err = decoder.Decode(&infoRedis)
	if err != nil {
		fmt.Println("Error parsing redisInfo:", err)
		return Config{}, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return Config{}, err
	}

	var infoAddr AddrInfo
	err = decoder.Decode(&infoAddr)
	if err != nil {
		fmt.Println("Error parsing addresses:", err)
		return Config{}, err
	}

	return Config{auth, db, brokersSlice, infoKafka, infoRedis, infoAddr}, nil
}
