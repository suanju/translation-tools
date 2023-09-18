package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func Cors() rest.RunOption {
	return rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization , beauthorizationarer")
		header.Set("Access-Control-Allow-Credentials", "true")
	}, func(w http.ResponseWriter) {
	}, "*")
}
