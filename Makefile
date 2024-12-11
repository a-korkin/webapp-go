run:
	go run cmd/main.go
prepare_db:
	chmod +x ./scripts/start.sh && ./scripts/start.sh
