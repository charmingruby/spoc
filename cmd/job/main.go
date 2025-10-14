package main

import (
	"github.com/charmingruby/spoc/internal/tm1"
	"github.com/charmingruby/spoc/internal/tm1/usecase"
)

func main() {
	data, err := tm1.NewFetcher().Fetch(usecase.FetchInput{
		APIKey:                      "2",
		ShouldSimulateAuthError:     false,
		ShouldSimulateRelatoryError: false,
	})

	if err != nil {
		println(err.Error())
	}

	println(string(data))
}
