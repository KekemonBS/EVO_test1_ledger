containers = postgres test1-ledger
up:
	docker compose up -d	
stop:
	docker compose stop
down:
	docker compose down
purge: down
	docker rmi -f $(containers)
rebuild:
	docker compose up --build -d
	docker image prune -a -f
