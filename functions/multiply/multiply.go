package multiply

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	conf "github.com/coding_challenge/pkg/config"
	"github.com/coding_challenge/pkg/errors"
	"github.com/coding_challenge/pkg/helper"

	"github.com/coding_challenge/pkg/model"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	functions.HTTP("Multiply", Multiply)

	cfg := &model.Config{}
	err := conf.GetConfig(cfg)
	if err != nil {
		log.Fatalf("%s Err: %s", errors.ErrGettingConfig, err.Error())
	}

	// Initializing redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: "",
		DB:       0,
	})
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, errors.ErrUnsupportedMethod, http.StatusMethodNotAllowed)
		return
	}

	resp, err := helper.ExtractQueryParams(r.URL.Query())
	if err != nil {
		http.Error(w, errors.ErrUnsupportedType, http.StatusBadRequest)
		return
	}

	key := fmt.Sprintf("%f*%f", resp.X, resp.Y)
	if resp.X > resp.Y {
		key = fmt.Sprintf("%f*%f", resp.Y, resp.X)
	}

	// Check if cached data exists
	result, err := redisClient.Get(key).Result()
	if err != nil && err != redis.Nil {
		http.Error(w, fmt.Sprintf("%s Err: %s", errors.ErrGettingCache, err.Error()), http.StatusInternalServerError)
		return
	}

	var jsonResp []byte
	var cachedResp *model.Response
	if result != "" {
		err = json.Unmarshal([]byte(result), &cachedResp)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s Err: %s", errors.ErrParsingCache, err.Error()), http.StatusInternalServerError)
			return
		}
		resp.Answer = cachedResp.Answer
		resp.Cached = true
	} else {
		resp.Answer = resp.X * resp.Y
		resp.Cached = false
	}

	resp.Action = "multiply"

	// Marshal response to json
	jsonResp, err = json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s Err: %s", errors.ErrJSONMarshal, err.Error()), http.StatusInternalServerError)
		return
	}

	//	Set redis key with ttl
	err = redisClient.Set(key, jsonResp, 1*time.Minute).Err()
	if err != nil {
		http.Error(w, fmt.Sprintf("%s Err: %s", errors.ErrSettingCache, err.Error()), http.StatusInternalServerError)
		return
	}

	// Write http response in json
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatalf("%s Err: %s", errors.ErrWrittingResponse, err.Error())
	}
	return
}
