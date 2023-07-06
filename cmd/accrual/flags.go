package main

import "flag"

type conf struct {
	dsn string
}

func parseFlags() conf {
	var config conf
	flag.StringVar(&config.dsn, "d", "postgresql://app:123qwe@localhost:5432/go-accrual?sslmode=disable", "postgres dsn")
	flag.Parse()
	return config
}
