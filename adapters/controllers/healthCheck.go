package controllers

import (
	"go-cla-practice/adapters/presenter"
	"net/http"
)

func HealthCheckController() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		presenter.Success(writer, "I'm fine\n")
	}
}

func HelloController() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonMap := map[string]string{
			"hello": "Hello",
		}
		presenter.Success(writer, jsonMap)
	}
}
