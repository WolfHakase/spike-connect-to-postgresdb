package main

import "testing"

// Test_integration_connectToDB_success is the integration test for connectToDB().
// If the test database is not running locally this test will fail.
// In a real project there would be a separate test database running that these types of tests connect to.
// The credentials for that test database would not be confidential, and could therefore be hardcoded in tests.
func Test_integration_connectToDB_success(t *testing.T) {
	db, _ := connectToDB(&DBConfig{
		host:     "localhost",
		port:     "5432",
		user:     "postgres",
		password: "docker",
		dbname:   "postgres",
	})

	if db == nil {
		t.Fatalf("db should not be nil")
	}
}

// Test_integration_connectToDB_fail is the integration test for connectToDB().
// This test verifies that the database returns an error with invalid credentials.
func Test_integration_connectToDB_fail(t *testing.T) {
	_, err := connectToDB(&DBConfig{
		host:     "localhost",
		port:     "5432",
		user:     "not_postgres",
		password: "not_docker",
		dbname:   "not_postgres",
	})

	if err == nil {
		t.Fatalf("err should not be nil")
	}
}

// Test_integration_connectToDB_noConfig is the integration test for connectToDB().
// This test verifies that the database returns an error with no config.
func Test_integration_connectToDB_noConfig(t *testing.T) {
	_, err := connectToDB(nil)

	if err != ErrorCanNotConnectDatabaseNoConfig {
		t.Fatalf("err should be ErrorCanNotConnectDatabaseNoConfig")
	}
}