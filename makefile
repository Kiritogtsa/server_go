GOCMD := go

GOBUILD := $(GOCMD) build 
GOCLEAN := $(GOCMD) clean


BINARY_NAME := meu_app 

all: startcompose

build:
	$(GOBUILD) -o $(BINARY_NAME) 

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOCMD) run main.go

startcompose:
	docker compose start 

stopcompose:
	docker compose stop

startdocker:
	docker start app 

stopdocker:
	docker stop app

docker-compose-build:
	docker compose up

docker-compose-rm:
	docker compose down

ps:
	docker ps