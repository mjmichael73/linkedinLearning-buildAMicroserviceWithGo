up:
	@docker compose up -d
down:
	@docker compose down
clean:
	@docker compose down --remove-orphans --volumes
	@docker compose up --build -d
ps:
	@docker compose ps
dbexec:
	@docker compose exec -it db bash
dbpsql:
	@docker compose exec -it db sh -c "psql -h localhost -U db_microservice_user -d db_microservice"
dbschema:
	@docker compose exec -it db sh -c "psql -h localhost -U db_microservice_user -d db_microservice -f /data/schema.sql"
dbseed:
	@docker compose exec -it db sh -c "psql -h localhost -U db_microservice_user -d db_microservice -f /data/seed.sql"
run:
	@go run main.go