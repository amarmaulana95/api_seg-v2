# api_seg
 
-Elo Rating Algorithm is widely used rating algorithm that is used to rank players in many competitive games. Players with higher ELO rating have a higher probability of winning a game than a player with lower ELO rating. After each game, ELO rating of players is updated. If a player with higher ELO rating wins, only a few points are transferred from the lower rated player. However if lower rated player wins, then transferred points from a higher rated player are far greater.

- Approach: P1: Probability of winning of player with rating2 P2: Probability of winning of player with rating1. P1 = (1.0 / (1.0 + pow(10, ((rating1 – rating2) / 400)))); P2 = (1.0 / (1.0 + pow(10, ((rating2 – rating1) / 400)))); Obviously, P1 + P2 = 1. The rating of player is updated using the formula given below :- rating1 = rating1 + K*(Actual Score – Expected score); In most of the games, “Actual Score” is either 0 or 1 means player either wins or loose. K is a constant. If K is of a lower value, then the rating is changed by a small fraction but if K is of a higher value, then the changes in the rating are significant. Different organizations set a different value of K.

- migration

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

-migrate create -ext sql -dir db/migrations create_transaction_table

-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose up

-migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/mydb?sslmode=disable" -verbose down


