# api_seg
 
- migration

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

-migrate create -ext sql -dir db/migrations create_transaction_table
-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose up
-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose down


