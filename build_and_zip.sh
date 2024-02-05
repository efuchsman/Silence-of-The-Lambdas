rm -rf silence_of_the_lambdas
rm -rf deployment.zip

GOOS=linux GOARCH=arm64 go build -o silence_of_the_lambdas silence_of_the_lambdas.go
zip deployment.zip silence_of_the_lambdas bootstrap .env
