package container

import (
	"fmt"
	"github.com/IvanKr8/zylo/internal/container/storage"
	"github.com/IvanKr8/zylo/internal/container/types"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

func CreateContainer(config types.ContainerConfig) (*types.Container, error) {
	id := generateID()
	containerPath := filepath.Join("/var/lib/zylo/containers", id)

	if err := os.MkdirAll(containerPath, 0755); err != nil {
		return nil, fmt.Errorf("error creating container directory: %v", err)
	}

	cmd := exec.Command(config.Command[0], config.Command[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("container startup error: %v", err)
	}

	container := &types.Container{
		ID:        id,
		Name:      config.Name,
		Status:    "Running",
		Pid:       cmd.Process.Pid,
		CreatedAt: time.Now(),
	}

	// Сохранение информации о контейнере
	if err := storage.SaveContainerInfo(container); err != nil {
		return nil, fmt.Errorf("error saving container information: %v", err)
	}

	return container, nil
}

func StartContainer(id string) error {
	container, err := storage.LoadContainerInfo(id)
	if err != nil {
		return fmt.Errorf("error loading container information: %v", err)
	}

	if container.Status == "Running" {
		return fmt.Errorf("the container is already running")
	}

	container.Status = "Running"
	return storage.SaveContainerInfo(container)
}

func StopContainer(id string) error {
	container, err := storage.LoadContainerInfo(id)
	if err != nil {
		return fmt.Errorf("error loading container information: %v", err)
	}

	if container.Status != "Running" {
		return fmt.Errorf("container is not running")
	}

	if err := syscall.Kill(container.Pid, syscall.SIGTERM); err != nil {
		return fmt.Errorf("container stop error: %v", err)
	}

	container.Status = "Stopped"
	return storage.SaveContainerInfo(container)
}

func RemoveContainer(id string) error {
	container, err := storage.LoadContainerInfo(id)
	if err != nil {
		return fmt.Errorf("error loading container information: %v", err)
	}

	if container.Status == "Running" {
		return fmt.Errorf("the container is running, it must be stopped before deleting")
	}

	containerPath := filepath.Join("/var/lib/zylo/containers", id)
	if err = os.RemoveAll(containerPath); err != nil {
		return fmt.Errorf("container deletion error: %v", err)
	}

	return storage.DeleteContainerInfo(id)
}

func GetContainerStatus(id string) (string, error) {
	container, err := storage.LoadContainerInfo(id)
	if err != nil {
		return "", fmt.Errorf("error loading container information: %v", err)
	}

	return container.Status, nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
