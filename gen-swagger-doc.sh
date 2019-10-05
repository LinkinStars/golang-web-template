#!/usr/bin/env bash
swag init -g ./cmd/main.go -o ./cmd/docs
swagger serve ./cmd/docs/swagger.yaml