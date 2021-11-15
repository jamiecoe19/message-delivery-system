# message-delivery-system
unity message delivery system

## Requirements
- golang
- Docker
- mysql

## My solution
- Use RabbitMQ to manage queues
- Use SQL to persist the users connected

## Endpoints
- /connect - connect user based on query param to hub
- /identity - send identiy message
- /list - list other users in hub
- /relay - send message to other user in hub
- /message - view message
- /disconnect - disconnect user based on query param from hub

## Run App
- `docker-compose up` to build rabbitmq and sql images
- go to location `cmd/mds` and run `go run main.go` 