postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root promo_news_apps

dropdb:
	docker exec -it postgres12 dropdb promo_news_apps

migrateup:
	migrate -path db/migration -database "postgresql://root:XiTvYhxjeNH9xSi8Bmg2@promo-news-apps.cdvqupkkttde.ap-southeast-1.rds.amazonaws.com:5432/promo_news_apps" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/promo_news_apps?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

