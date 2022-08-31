package env

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

type LoaderMoc struct {
	Data string
}

func NewLoadMoc(data string) ILoader {
	return &LoaderMoc{Data: data}
}

func (l *LoaderMoc) LoadFile(fileName string) (io.Reader, error) {
	var buffer bytes.Buffer
	buffer.WriteString(l.Data)
	return &buffer, nil
}

func TestFileNameOrDefault(t *testing.T) {
	assert.Equal(t, fileNameOrDefault(""), ".env")
	assert.Equal(t, fileNameOrDefault(".env"), ".env")
	assert.Equal(t, fileNameOrDefault("./../.env"), "./../.env")
}

func TestGetEnvMap(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("DATABASE_NAME=test\nNAME=test name\n")
	expected := make(map[string]string)
	expected["DATABASE_NAME"] = "test"
	expected["NAME"] = "test name"
	content, err := getEnvMap(&buffer)
	if err != nil {
		t.Error("Failed to get env map")
	}
	assert.Equal(t, content, expected)
}

func TestSetToEnv(t *testing.T) {
	os.Clearenv()
	envMap := make(map[string]string)
	envMap["DATABASE_NAME"] = "test"
	envMap["NAME"] = "test name"
	err := setToEnv(envMap)
	assert.Equal(t, err, nil)
	assert.Equal(t, os.Getenv("DATABASE_NAME"), "test")
	assert.Equal(t, os.Getenv("NAME"), "test name")

	os.Clearenv()
	envMap = make(map[string]string)
	envMap["=="] = "test"
	err = setToEnv(envMap)
	fmt.Println(err)
	assert.NotNil(t, err)
}

func TestLoad(t *testing.T) {
	os.Clearenv()
	err := LoadFile{LoadRepo: NewLoadMoc("DATABASE_NAME=test\nDATABASE_PORT=9876\n")}.Load("")
	assert.Equal(t, os.Getenv("DATABASE_NAME"), "test")
	assert.Equal(t, os.Getenv("DATABASE_PORT"), "9876")
	assert.Equal(t, err, nil)

	os.Clearenv()
	err = LoadFile{LoadRepo: NewLoadMoc(`==test\n`)}.Load("")
	assert.NotNil(t, err)
}
