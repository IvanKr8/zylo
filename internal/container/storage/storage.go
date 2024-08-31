package storage

import (
	"encoding/json"
	"fmt"
	"github.com/IvanKr8/zylo/internal/container/types"
	"os"
	"path/filepath"
)

func SaveContainerInfo(container *types.Container) error {
	data, err := json.Marshal(container)
	if err != nil {
		return fmt.Errorf("ошибка сериализации данных контейнера: %v", err)
	}

	containerFile := filepath.Join("/var/lib/zylo/containers", container.ID, "config.json")
	if err := os.WriteFile(containerFile, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи данных контейнера: %v", err)
	}

	return nil
}

func LoadContainerInfo(id string) (types.Container, error) {
	containerFile := filepath.Join("/var/lib/zylo/containers", id, "config.json")
	data, err := os.ReadFile(containerFile)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения данных контейнера: %v", err)
	}

	var container types.Container
	if err = json.Unmarshal(data, &container); err != nil {
		return nil, fmt.Errorf("ошибка десериализации данных контейнера: %v", err)
	}

	return &container, nil
}

func DeleteContainerInfo(id string) error {
	containerFile := filepath.Join("/var/lib/zylo/containers", id)
	return os.RemoveAll(containerFile)
}
