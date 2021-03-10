package main

import (
	zconfig "ze-redis-test/config"
)

// Currently HARDCODED
// TODO: implement normal config reading from:
//  ENV variable
//  conf file
// STUB:
func getConfig() (zconfig.Conf, error) {

	// fullfill appParams from config
	// HARDCODE it temporary
	appParams := map[string]string{
		"db_type":     "redis",
		"db_addr":     "127.0.0.1",
		"db_port":     "6379",
		"db_password": "",
		"db_num":      "0",
		"app_port":    "8421",
	}

	// not implemented yet
	// validateParams()

	return zconfig.Conf{Params: appParams}, nil
}
