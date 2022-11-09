# api_seg
 
- SEG Is an Innovation system that can recommend the most accurate construction costs. This application already uses an Algorithm to be able to recommend.
With SEG you get a wide variety of methods, types of estimation, innovation, value engineering, financial engineering.

- Ex : P1: Probability of winning of player with rating2 P2: Probability of winning of player with rating1. P1 = (1.0 / (1.0 + pow(10, ((rating1 – rating2) / 400)))); P2 = (1.0 / (1.0 + pow(10, ((rating2 – rating1) / 400)))); Obviously, P1 + P2 = 1. The rating of player is updated using the formula given below :- rating1 = rating1 + K*(Actual Score – Expected score); In most of the games, “Actual Score” is either 0 or 1 means player either wins or loose. K is a constant. If K is of a lower value, then the rating is changed by a small fraction but if K is of a higher value, then the changes in the rating are significant. Different organizations set a different value of K.

- migration

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

-migrate create -ext sql -dir db/migrations create_transaction_table

-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose up

-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose down


