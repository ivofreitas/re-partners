package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"sync"
)

// Env values
type Env struct {
	Server Server
}

// Server config
type Server struct {
	Port      string
	PackSizes []int
}

var (
	env  *Env
	once sync.Once
)

// GetEnv returns env values
func GetEnv() *Env {

	once.Do(func() {

		viper.AutomaticEnv()
		godotenv.Load("./internal/config/.env")

		env = new(Env)
		env.Server.Port = viper.GetString("PORT")
		packSizesStr := viper.GetString("PACK_SIZES")
		packSizesArray := strings.Split(packSizesStr, ",")
		for _, sizeStr := range packSizesArray {
			size, err := strconv.Atoi(strings.TrimSpace(sizeStr))
			if err != nil {
				panic(fmt.Errorf("invalid pack size: %s", sizeStr))
			}
			env.Server.PackSizes = append(env.Server.PackSizes, size)
		}
	})
	return env
}
