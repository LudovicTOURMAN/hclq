package pkg

import (
	"bytes"
	"io"
	"os"

	hclq "github.com/LudovicTOURMAN/hclq/cmd"
)

// SetTerraformAttribute Set attribute value
func SetTerraformAttribute(tfParamPath, tfMainLocation, value string) (err error) {
	if _, err = os.Stat(tfMainLocation); os.IsNotExist(err) {
		return
	}

	cmd := hclq.RootCmd
	cmd.SetArgs([]string{
		"-i", tfMainLocation,
		"-o", tfMainLocation,
		"set", tfParamPath, string(value),
	})
	return cmd.Execute()
}

// GetTerraformAttribute Get attribute value
func GetTerraformAttribute(tfParamPath, tfMainLocation string) (result string, err error) {
	if _, err := os.Stat(tfMainLocation); os.IsNotExist(err) {
		return "", err
	}

	cmd := hclq.RootCmd
	cmd.SetArgs([]string{
		"-i", tfMainLocation,
		"get", tfParamPath,
	})
	result, err = CaptureStdout(cmd.Execute)

	return
}

// CaptureStdout Record stdout and retrieve it as a string
func CaptureStdout(f func() error) (stdout string, err error) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if err = f(); err != nil {
		return
	}

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	stdout = buf.String()
	return
}
