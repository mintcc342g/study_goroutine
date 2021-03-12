# study_goroutine

* Go: 1.14

```
# build
(~/project_root) $ make build

# run server use command line
(~/project_root) $ go run ./api/main.go

# OR run server as debug mode using vscode (click the button of 'Launch')

# install docker if you cloned v1.0.0
# then, run docker (you can change the image which you want)
$ docker run -d --hostname studymq --name studymq -p 5672:5672 --restart=unless-stopped -e RABBITMQ_DEFAULT_USER=mquser -e RABBITMQ_DEFAULT_PASS=mqpwd rabbitmq:3

## How to Test

# Normal response test
$curl --location --request POST '127.0.0.1:1323/api/v1/email/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "senderAddress": "sender@email.com",
    "receiverAddress": "receiver@email.com",
    "title": "anything",
    "content": "normal case"
}'

# Infinite Loop test
$curl --location --request POST '127.0.0.1:1323/api/v1/email/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "senderAddress": "sender@email.com",
    "receiverAddress": "receiver@email.com",
    "title": "err",
    "content": "this would raise an error."
}'
```
