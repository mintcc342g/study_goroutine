# study_goroutine

* Go: 1.14

```
# build
(~/project_root) $ make build

# run server use command line
(~/project_root) $ go run ./api/main.go

# OR run server as debug mode using vscode (click the button of 'Launch')


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
