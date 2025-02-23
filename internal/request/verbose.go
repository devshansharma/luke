package request

import (
	"net/http"

	"github.com/devshansharma/luke/pkg/writer"
)

func writeRequest(req *http.Request, cfg RequestConfig) {
	if cfg.Silent {
		return
	}

	writer := writer.GetInstance()

	if cfg.Verbose {
		writer.Log("Request Headers:")
		for key, values := range req.Header {
			for _, val := range values {
				writer.Log("%s: %s", key, val)
			}
		}
		writer.Log("")

		if cfg.Method == "POST" || cfg.Method == "PUT" || cfg.Method == "PATCH" {
			writer.Log("Request Body:")
			writer.Log(cfg.Data)
			writer.Log("")
		}
	}
}

func writeResponse(response *http.Response, body []byte, cfg RequestConfig) {
	writer := writer.GetInstance()

	if !cfg.Silent && cfg.Verbose {
		writer.Response("Response Headers:")
		for key, values := range response.Header {
			for _, val := range values {
				writer.Response("%s: %s", key, val)
			}
		}
		writer.Log("")
		writer.Log("Response Body:")
	}

	writer.Response(string(body))
}
