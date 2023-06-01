# plist api

# Run api through the following steps

# Dockerfile

# create a local .env file in the working directory and define the port number
# for ex. APP_PORT=54321
docker build . -t plist
docker run plist

# Docker Compose

# create a local .env file in the working directory and define the port number
# for ex. APP_PORT=54321
docker compose up -d

# Golang Executable
make appserver
./appserver 65333

# Postman example request
0.0.0.0:65333/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z

# Run tests
make test