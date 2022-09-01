package env

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type ILoader interface {
	LoadFile(fileName string) (io.Reader, error)
}

type Loader struct {
}

func NewLoader() ILoader {
	return &Loader{}
}

type LoadFile struct {
	LoadRepo ILoader
}

// fileNameOrDefault returns file name if there is no file name then returns default ".env"
func fileNameOrDefault(fileName string) string {
	if len(fileName) > 0 {
		return fileName
	}
	return ".env"
}

// getEnvMap gets data from environment
func getEnvMap(file io.Reader) (map[string]string, error) {
	envMap := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitString := strings.Split(line, "=")
		envMap[splitString[0]] = splitString[1]
	}
	return envMap, nil
}

// setToEnv sets data in environment
func setToEnv(envMap map[string]string) error {
	for key, value := range envMap {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadFile load file form file name
func (l *Loader) LoadFile(fileName string) (io.Reader, error) {
	return os.Open(fileName)
}

// Load is loads data from file and set into env
func (l LoadFile) Load(fileName string) error {
	var envMap map[string]string
	fileName = fileNameOrDefault(fileName)
	file, err := l.LoadRepo.LoadFile(fileName)
	if err != nil {
		return err
	}
	envMap, err = getEnvMap(file)
	if err != nil {
		return err
	}
	err = setToEnv(envMap)
	return err
}
