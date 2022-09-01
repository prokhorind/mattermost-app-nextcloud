package main

import (
	"github.com/prokhorind/nextcloud/function/install"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prokhorind/nextcloud/function"
)

func main() {
	r := gin.Default()

	r.StaticFS("/static/", http.Dir(os.Getenv("STATIC_FOLDER")))

	function.InitHandlers(r)

	r.GET("/manifest.json", install.GetManifest)

	port := getPort()

	r.Run(":" + port)

}

func getPort() string {

	portStr := os.Getenv("PORT")

	if portStr == "" {
		u, err := url.Parse(getEnv("APP_URL", "http://localhost:8082"))
		if err != nil {
			panic(err)
		}
		portStr = u.Port()
	}

	return portStr
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}
