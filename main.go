package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/LoyalPotato/commit-helper/src/cmd"
)

func main() {
	loadEnv()
	if cmdErr := cmd.Execute(); cmdErr != nil {
		fmt.Println(cmdErr)
		os.Exit(1)
	}
}

func loadEnv() {
	envs := []string{".env", ".env.local", ".local.env"}

	for _, env := range envs {
		godotenv.Load(env)
	}
}
