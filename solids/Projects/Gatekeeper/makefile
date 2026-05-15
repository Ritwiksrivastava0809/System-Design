docker-network:
	docker network create gatekeeper-network || true

compose-up: docker-network
	docker-compose up -d

server:
	go run cmd/api/main.go -e development



.PHONY: docker-network , compose-up, server