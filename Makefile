default: build_server build_client

server:
	go build -o bin/PiDoorIntercom cmd/PiDoorIntercom/main.go

webui:
	cd web && npm run build