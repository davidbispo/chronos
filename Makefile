DB_DSN=mysql://chronos:chronos@tcp(db:3306)/chronos_db
MIGRATIONS_DIR=db/migrations

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" down 1

migrate-drop:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" drop -f

migrate-new:
ifndef name
	$(error You must provide a name, example: make migrate-new name=create_users_table)
endif
	migrate create -ext sql -dir $(MIGRATIONS_DIR) $(name)

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_DSN)" version