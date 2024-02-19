package configs

import (
	"strconv"
	"testing"
)

func TestParsedConfig(t *testing.T) {
	// arrange
	expectedUsername := "USERNAME"
	expectedPassword := "PASSWORD"
	expectedServerPort := 80
	t.Setenv("DB_USERNAME", expectedUsername)
	t.Setenv("DB_PASSWORD", expectedPassword)
	t.Setenv("SERVER_PORT", strconv.Itoa(expectedServerPort))

	// act
	config, err := ParsedConfig()

	// assert
	if err != nil {
		t.Fatalf("failed with error %s", err)
	}

	if config.Database.USERNAME != expectedUsername ||
		config.Database.PASSWORD != expectedPassword ||
		config.ServerPort != expectedServerPort {
		t.Fatalf(`
			Received USERNAME: %s, PASSWORD: %s, SERVERPORT: %d,
			Expected USERNAME: %s, PASSWORD: %s, SERVERPORT: %d
		`, config.Database.USERNAME, config.Database.PASSWORD, config.ServerPort,
			expectedUsername, expectedPassword, expectedServerPort)
	}
}
