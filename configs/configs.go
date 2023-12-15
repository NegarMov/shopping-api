package configs

import (
    "encoding/json"
    "os"
)

type Config struct {
	ServerIP	string	`json:"server_ip"`
    ServerPort	uint	`json:"server_port"`
    DBHost 		string	`json:"db_host"`
    DBUser 		string	`json:"db_user"`
    DBPass 		string	`json:"db_pass"`
    DBName 		string	`json:"db_name"`
    DBPort 		uint	`json:"db_port"`
}

func LoadConfig() (Config, error) {
    var config Config

    file, err := os.Open("./configs/configs.json")
    if err != nil {
        return config, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    if err != nil {
        return config, err
    }

    return config, nil
}
