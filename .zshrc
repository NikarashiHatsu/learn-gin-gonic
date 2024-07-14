## Migrations
### Making Migrations
alias go-make-migration='migrate create -ext sql -dir database/migrations -seq'

### MySQL
alias go-migrate-mysql='migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/go_gin_gonic" -source database/migrations up'
alias go-migrate-rollback-mysql='migrate -database "mysql://root:root@tcp(127.0.0.1:3306)/go_gin_gonic" -source database/migrations down'

### PostgreSQL
alias go-migrate-psql='migrate -database "postgres://root:root@127.0.0.1:5432/go_gin_gonic" -source database/migrations up'
alias go-migrate-rollback-psql='migrate -database "postgres://root:root@127.0.0.1:5432/go_gin_gonic" -source database/migrations down'

### Default Config
alias go-migrate='go-migrate-mysql'
alias go-migrate-rollback='go-migrate-rollback-mysql'