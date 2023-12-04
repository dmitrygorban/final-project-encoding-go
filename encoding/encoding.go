package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	bytes, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(bytes, &j.DockerCompose); err != nil {
		return err
	}

	yamlConfig, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}

	f, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(yamlConfig)
	if err != nil {
		return err
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	bytes, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(bytes, &y.DockerCompose); err != nil {
		return err
	}

	jsonConfig, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}

	f, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(jsonConfig)
	if err != nil {
		return err
	}

	return nil

}
