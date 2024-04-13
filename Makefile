DB_URL := "postgres://postgres:nodirbek@localhost:5432/postgres?sslmode=disable"

migrate-up:
    migrate -path ./db/migrations -database "$(DB_URL)" -verbose up

migrate-down:
    migrate -path ./db/migrations -database "$(DB_URL)" -verbose down

migrate-file:
    migrate create -ext sql -dir db/migrations/ -seq create_users_table
