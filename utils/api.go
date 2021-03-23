package utils

import (
	coingecko "github.com/superoo7/go-gecko/v3"
	"net/http"
	"time"
)

var (
	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}
	API = coingecko.NewClient(httpClient)
)
