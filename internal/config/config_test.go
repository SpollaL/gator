package config

import (
	"testing"
)

func TestConfigRead(t *testing.T) {
	config, err := Read()
	if err != nil {
		t.Errorf("Reading config terminated with %s", err)
	}

	if config.DbUrl != "postgres://example" {
		t.Error("The Dburl should have been postgres://example")
	}
}

func TestConfigSetUser(t *testing.T) {
	config, err := Read()
	if err != nil {
		t.Errorf("Reading config terminated with %s", err)
	}
	config.SetUser("test_user")
	config_new, err := Read()
	if err != nil {
		t.Errorf("Reading config terminated with %s", err)
	}
	if config_new.CurrentUserName != "test_user" {
		t.Error("Current user should be equal to test user")
	}
}
