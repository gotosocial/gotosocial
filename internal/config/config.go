/*
   GoToSocial
   Copyright (C) 2021 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Config pulls together all the configuration needed to run gotosocial
type Config struct {
	LogLevel        string    `yaml:"logLevel"`
	ApplicationName string    `yaml:"applicationName"`
	DBConfig        *DBConfig `yaml:"db"`
}

// New returns a new config, or an error if something goes amiss.
// The path parameter is optional, for loading a configuration json from the given path.
func New(path string) (*Config, error) {
	config := &Config{
		DBConfig: &DBConfig{},
	}
	if path != "" {
		var err error
		if config, err = loadFromFile(path); err != nil {
			return nil, fmt.Errorf("error creating config: %s", err)
		}
	}

	return config, nil
}

// loadFromFile takes a path to a yaml file and attempts to load a Config object from it
func loadFromFile(path string) (*Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read file at path %s: %s", path, err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("could not unmarshal file at path %s: %s", path, err)
	}

	return config, nil
}

// ParseFlags sets flags on the config using the provided Flags object
func (c *Config) ParseFlags(f KeyedFlags) {
	fn := GetFlagNames()

	// For all of these flags, we only want to set them on the config if:
	//
	// a) They haven't been set at all in the config file we already parsed,
	// 	  and so we take the default from the flags object.
	//
	// b) They may have been set in the config, but they've *also* been set explicitly
	//    as a command-line argument or an env variable, which takes priority.

	// general flags
	if c.LogLevel == "" || f.IsSet(fn.LogLevel) {
		c.LogLevel = f.String(fn.LogLevel)
	}

	if c.ApplicationName == "" || f.IsSet(fn.ApplicationName) {
		c.ApplicationName = f.String(fn.ApplicationName)
	}

	// db flags
	if c.DBConfig.Type == "" || f.IsSet(fn.DbType) {
		c.DBConfig.Type = f.String(fn.DbType)
	}

	if c.DBConfig.Address == "" || f.IsSet(fn.DbAddress) {
		c.DBConfig.Address = f.String(fn.DbAddress)
	}

	if c.DBConfig.Port == 0 || f.IsSet(fn.DbPort) {
		c.DBConfig.Port = f.Int(fn.DbPort)
	}

	if c.DBConfig.User == "" || f.IsSet(fn.DbUser) {
		c.DBConfig.User = f.String(fn.DbUser)
	}

	if c.DBConfig.Password == "" || f.IsSet(fn.DbPassword) {
		c.DBConfig.Password = f.String(fn.DbPassword)
	}

	if c.DBConfig.Database == "" || f.IsSet(fn.DbDatabase) {
		c.DBConfig.Database = f.String(fn.DbDatabase)
	}
}

// KeyedFlags is a wrapper for any type that can store keyed flags and give them back.
// HINT: This works with a urfave cli context struct ;)
type KeyedFlags interface {
	String(k string) string
	Int(k string) int
	IsSet(k string) bool
}

// Flags is used for storing the names of the various flags used for
// initializing and storing urfavecli flag variables.
type Flags struct {
	LogLevel        string
	ApplicationName string
	ConfigPath      string
	DbType          string
	DbAddress       string
	DbPort          string
	DbUser          string
	DbPassword      string
	DbDatabase      string
}

// GetFlagNames returns a struct containing the names of the various flags used for
// initializing and storing urfavecli flag variables.
func GetFlagNames() Flags {
	return Flags{
		LogLevel:        "log-level",
		ApplicationName: "application-name",
		ConfigPath:      "config-path",
		DbType:          "db-type",
		DbAddress:       "db-address",
		DbPort:          "db-port",
		DbUser:          "db-user",
		DbPassword:      "db-password",
		DbDatabase:      "db-database",
	}
}

// GetEnvNames returns a struct containing the names of the environment variable keys used for
// initializing and storing urfavecli flag variables.
func GetEnvNames() Flags {
	return Flags{
		LogLevel:        "GTS_LOG_LEVEL",
		ApplicationName: "GTS_APPLICATION_NAME",
		ConfigPath:      "GTS_CONFIG_PATH",
		DbType:          "GTS_DB_TYPE",
		DbAddress:       "GTS_DB_ADDRESS",
		DbPort:          "GTS_DB_PORT",
		DbUser:          "GTS_DB_USER",
		DbPassword:      "GTS_DB_PASSWORD",
		DbDatabase:      "GTS_DB_DATABASE",
	}
}
