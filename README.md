# message-delivery-system
unity message delivery system

## Requirements
- golang
- Docker
- mysql

## My solution
- Use RabbitMQ to manage queues
- Use SQL to persist the users connected
- Use Http echo to connect and discount users from the system and send messages

## Endpoints
-  GET /connect?name={name} - connect user based on query param to hub
-  GET /identity?name={name} - send identiy message
-  GET /list?name={name} - list other users in hub
-  POST /relay - send message to other user in hub
-  DELETE /disconnect?name={name} - disconnect user based on query param from hub

## Run App
- `docker-compose up` to build rabbitmq and sql images
- add migration script in `cmd/mds/migration-script` and run `schema.sql`
- go to location `cmd/mds` and run `go run main.go`
    - all messages sent will be displayed in the console with the assoicated user name and rabbitmq queue id

## Run e2e tests
- `docker-compose up` to build rabbitmq and sql images
- add migration script in `cmd/mds/migration-script` and run `schema.sql`
- run package `e2e` in `cmd` directory