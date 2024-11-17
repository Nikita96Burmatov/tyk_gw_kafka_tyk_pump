ifneq (,$(wildcard ./.env))
    include .env
    export
endif

up:
	docker-compose --env-file .env up -d

down:
	docker-compose --env-file .env down

restart:
	make down && make up
