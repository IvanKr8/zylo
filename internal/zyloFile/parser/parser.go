package parser

import (
	"bufio"
	"fmt"
	"github.com/IvanKr8/zylo/internal/zyloFile/types"
	"io/ioutil"
	"strings"
)

func ZyloParser(zyloFile string) (*types.ZyloConfig, error) {
	content, err := ioutil.ReadFile(zyloFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	config := &types.ZyloConfig{
		EnvVars: make(map[string]string),
		Copies:  make(map[string]string),
	}

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "!!") {
			continue
		}

		if err = parseLine(line, config); err != nil {
			return nil, fmt.Errorf("failed to parse line: %v", err)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %v", err)
	}

	return config, nil
}

func parseLine(line string, config *types.ZyloConfig) error {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return fmt.Errorf("invalid command: %s", line)
	}

	switch parts[0] {
	case "USE_IMAGE":
		return handleUseImage(parts[1:], config)
	case "SET_WORKDIR":
		return handleSetWorkdir(parts[1:], config)
	case "COPY_TO":
		return handleCopyTo(parts[1:], config)
	case "EXECUTE":
		return handleExecute(parts[1:], config)
	case "SET_ENV":
		return handleSetEnv(parts[1:], config)
	case "OPEN_PORT":
		return handleOpenPort(parts[1:], config)
	case "START_WITH":
		return handleStartWith(parts[1:], config)
	default:
		return fmt.Errorf("unknown command: %s", parts[0])
	}
}
