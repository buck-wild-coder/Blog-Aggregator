package main

import (
	"fmt"

	"github.com/buck/blog/internal/config"
)

func main() {
	data, err := config.Read()
	if err != nil {
		return
	}

	cfg := config.Config{
		DBURL: data.DBURL,
	}

	cfg.SetUser("lane")
	data, err = config.Read()
	if err != nil {
		return
	}

	fmt.Println(data)
}
