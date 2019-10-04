#!/usr/bin/env bash
swag init -g ./src/cmd/main.go -o ./src/cmd/docs
swagger serve ./src/cmd/docs/swagger.yaml