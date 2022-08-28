package env

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func fileNameOrDefault(fileName string) string {
	if len(fileName) > 0 {
		return fileName
	}
	return ".env"
}

func getEnvMap(fileName string) (map[string]string, error) {
	envMap := make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitString := strings.Split(line, "=")
		envMap[splitString[0]] = splitString[1]
	}
	return envMap, nil
}

func Load(fileName string) error {
	fileName = fileNameOrDefault(fileName)
	envMap, err := getEnvMap(fileName)
	if err != nil {
		return err
	}
	for key, value := range envMap {
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}
