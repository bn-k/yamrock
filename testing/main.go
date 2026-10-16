package main

import (
	"fmt"
	"github.com/bn-k/yamrock/testing/config"
)

func main() {
	conf := config.NewConfig()

	fmt.Println(conf.Server().Address())
}
