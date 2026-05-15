docker-network:
	docker network create gatekeeper-network || true

compose-up: docker-network
	docker-compose up -d
	
.PHONY: docker-network