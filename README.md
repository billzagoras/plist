# plist api

# Run api through the following steps

# 1. Docker Compose

# define the port number in .env file, for ex. APP_PORT=54321
docker compose up -d

# 2. Golang Executable
make appserver
./appserver 65333

# Postman example request
0.0.0.0:65333/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z

# Run tests
make test
