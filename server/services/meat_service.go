package services

import (
    "strings"
)

type MeatService struct{}
type Result struct {
    Message string `json:"message"`
}

func (_ MeatService) Call(message string) (Result) {
    message = strings.TrimPrefix(message, "/")
    return Result{Message: message}
}
