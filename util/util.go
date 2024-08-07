package util

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type request struct{}

func LoadEnviromentalVariables(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(".env file is missing")
	}
	return os.Getenv(key)
}

func DecodeListReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request
	return req, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
