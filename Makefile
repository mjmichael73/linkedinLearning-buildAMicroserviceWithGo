down:
	@docker compose down --remove-orphans --volumes
build:
	@docker compose build
up: down build
	@docker compose down --remove-orphans --volumes
	@docker compose up --build -d
	@sleep 5
	@docker compose exec -it db sh -c "psql -h localhost -U wisdom_db_user -d wisdom_db -f /data/schema.sql"
	@docker compose exec -it db sh -c "psql -h localhost -U wisdom_db_user -d wisdom_db -f /data/seed.sql"
	@docker compose up -d
ps:
	@docker compose ps