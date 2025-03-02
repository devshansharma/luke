package request

import (
	"fmt"
	"reflect"
	"time"

	"github.com/spf13/cobra"
)

type Request struct {
	ID          string   `json:"id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"omitempty"`
	URL         string   `json:"url" validate:"required,url"`
	Method      string   `json:"method" validate:"oneof=GET POST PUT PATCH DELETE"`
	Headers     []string `json:"headers" validate:"omitempty"`
	DataType    string   `json:"data_type" validate:"omitempty"`
	Data        string   `json:"data" validate:"omitempty"`
	PreScript   string   `json:"pre_script" validate:"omitempty"`
	PostScript  string   `json:"post_script" validate:"omitempty"`
}

// RequestConfig to be used when we use it as a curl alternative
type RequestConfig struct {
	URL string `validate:"required,url"`

	Method  string        `flag:"method" validate:"oneof=GET POST PUT PATCH DELETE"`
	Headers []string      `flag:"header" validate:"omitempty"`
	Data    string        `flag:"data" validate:"omitempty"`
	Timeout time.Duration `flag:"timeout" validate:"omitempty"`

	Output  string `flag:"output" validate:"omitempty"`
	LogFile string `flag:"log-file" validate:"omitempty"`
	Silent  bool   `flag:"silent" validate:"omitempty"`
	Verbose bool   `flag:"verbose" validate:"omitempty"`
}

func ParseFlags(cmd *cobra.Command, config any) (err error) {
	flags := cmd.Flags()

	val := reflect.ValueOf(config).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		flagName := typ.Field(i).Tag.Get("flag")

		if flagName == "" {
			continue // Skip fields without flag tags
		}

		switch field.Kind() {
		case reflect.String:
			value, _ := flags.GetString(flagName)
			field.SetString(value)

		case reflect.Bool:
			value, _ := flags.GetBool(flagName)
			field.SetBool(value)

		case reflect.Int64:
			value, _ := flags.GetInt64(flagName)

			if typ.Field(i).Type == reflect.TypeOf(time.Duration(0)) {
				field.Set(reflect.ValueOf(time.Duration(value) * time.Millisecond))
			} else {
				field.SetInt(value)
			}

		case reflect.Slice:
			value, _ := flags.GetStringSlice(flagName)
			field.Set(reflect.ValueOf(value))

		default:
			return fmt.Errorf("unsupported field type: %s", field.Kind())
		}
	}

	return
}
