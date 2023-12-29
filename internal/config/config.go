package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error{
	err := godotenv.Load("../../.env")
	if err != nil {
		return fmt.Errorf("load env data is not working properly %v", err)
	}

	return nil
}