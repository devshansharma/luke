package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/devshansharma/luke/pkg/writer"
)

// HandleRequest for handling request
func HandleRequest(cfg RequestConfig) (err error) {
	req, cancel, err := createHttpRequest(cfg)
	if err != nil {
		return
	}
	defer cancel()

	client := http.Client{
		Timeout: cfg.Timeout,
	}

	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var body bytes.Buffer
	_, err = io.Copy(&body, response.Body)
	if err != nil {
		return
	}

	writer := writer.GetInstance()

	if cfg.Verbose {
		writer.Log(req.Method + " " + req.URL.String())
	}

	if !cfg.Silent {
		writer.Log(response.Status)
	}

	writeRequest(req, cfg)
	writeResponse(response, body.Bytes(), cfg)

	return
}

func createHttpRequest(cfg RequestConfig) (*http.Request, context.CancelFunc, error) {
	var cancel context.CancelFunc
	ctx := context.Background()
	if cfg.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, cfg.Timeout)
	}

	body, err := getRequestPayload(cfg)
	if err != nil {
		return nil, cancel, err
	}

	req, err := http.NewRequestWithContext(ctx, cfg.Method, cfg.URL, body)
	if err != nil {
		return nil, cancel, fmt.Errorf("failed to create request: %s", err)
	}

	return req, cancel, nil
}

func getRequestPayload(cfg RequestConfig) (body io.Reader, err error) {
	var (
		payload bytes.Buffer
		data    []byte
	)

	if cfg.Method == "GET" {
		body = http.NoBody
		return
	}

	if cfg.Data != "" {
		if strings.HasPrefix(cfg.Data, "@") {
			data, err = os.ReadFile(cfg.Data[1:])
			if err != nil {
				err = fmt.Errorf("failed to read file: %s", err)
				return
			}

			payload.Write(data)
		} else {
			payload.WriteString(cfg.Data)
		}

		body = &payload
	}

	return
}
