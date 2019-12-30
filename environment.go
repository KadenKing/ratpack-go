package main

import "os"

type environment interface {
	Get(key string) string
}

type env struct{}

func newEnv() *env {
	return &env{}
}

func (e *env) Get(key string) string {
	return os.Getenv(key)
}
