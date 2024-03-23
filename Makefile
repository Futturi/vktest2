sqlup:
	migrate -source file://internal/migrate -database postgres://root:root@localhost:5432/root?sslmode=disable up
sqldown:
	migrate -source file://internal/migrate -database postgres://root:root@localhost:5432/root?sslmode=disable down
.PHONY: sqlup, sqldown