back:
	SERVER_PORT=3302 go run server/main.go

front:
	yarn
	PORT=3002 yarn start
