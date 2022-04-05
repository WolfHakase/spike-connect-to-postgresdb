package main

import (
	"os"
	"testing"
)

func Test_integration_getDatabaseConfig(t *testing.T) {
	_ = loadGoDotEnv()
	config := getDatabaseConfig()

	expectedConfig := DBConfig{
		host:     "localhost",
		port:     "5432",
		user:     "postgres",
		password: "docker",
		dbname:   "postgres",
	}

	if *config != expectedConfig {
		t.Fatalf("config is not equal to expected")
	}
}

func Test_integration_getDatabaseConfig_notloaded(t *testing.T) {
	config := getDatabaseConfig()

	expectedConfig := DBConfig{
		host:     "",
		port:     "",
		user:     "",
		password: "",
		dbname:   "",
	}

	if *config != expectedConfig {
		t.Fatalf("config is not equal to expected")
	}
}

func Test_integration_getDatabaseConfig_overloaded(t *testing.T) {
	os.Setenv("DB_HOST", "not_localhost")
	os.Setenv("DB_PORT", "not_5432")
	os.Setenv("DB_USERNAME", "not_postgres")
	os.Setenv("DB_PASSWORD", "not_docker")
	os.Setenv("DB_NAME", "not_postgres")

	config := getDatabaseConfig()

	expectedConfig := DBConfig{
		host:     "not_localhost",
		port:     "not_5432",
		user:     "not_postgres",
		password: "not_docker",
		dbname:   "not_postgres",
	}

	if *config != expectedConfig {
		t.Fatalf("config is not equal to expected")
	}
}
