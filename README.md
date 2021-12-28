# Study Goroutine

* Go: 1.14

<br/>

## Build

```bash
# run command on project root directory
$ make build

```

<br/>

## Run Server

```bash
# run RabbitMQ using Docker
$ docker run -d --hostname emailmq --name emailmq -p 5672:5672 --restart=unless-stopped -e RABBITMQ_DEFAULT_USER=mquser -e RABBITMQ_DEFAULT_PASS=mqpwd rabbitmq:3

# run server (You can use debug mode on vscode as clicking the button of 'Launch')
$ go run ./api/main.go
```

<br/>

## How To Test

```bash
# Normal Case
$curl --location --request POST '127.0.0.1:1323/api/v1/email/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "senderAddress": "sender@email.com",
    "receiverAddress": "receiver@email.com",
    "title": "anything",
    "content": "normal case"
}'

# Error Case (Infinite Loop)
$curl --location --request POST '127.0.0.1:1323/api/v1/email/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "senderAddress": "sender@email.com",
    "receiverAddress": "receiver@email.com",
    "title": "err",
    "content": "this would raise an error."
}'
```