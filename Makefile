include .env
migrate-up:
	@migrate -source file://./_sql -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}' up
	#@migrate -source file://./_sql -database 'mysql://demo_user:demopass@tcp(db:3306)/demo' up
	# テスト用 @migrate -source file://./_sql -database 'mysql://$(TEST_DB_USER):$(TEST_DB_PASS)@tcp($(DB_HOST):3306)/$(TEST_DB_NAME)' up
migrate-down:
	@migrate -source file://./_sql -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:3306)/${DB_NAME}' down
	# テスト用 @migrate -source file://./_sql -database 'mysql://${TEST_DB_USER}:${TEST_DB_PASS}@tcp(${DB_HOST}:3306)/${TEST_DB_NAME}' down