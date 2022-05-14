test :
	./bin/scoreboard -path=./scoreboard.db

build :
	go build -o=./bin/scoreboard ./cmd/scoreboard

run :
	go run ./cmd/scoreboard -path=./scoreboard.db

install : ./bin/scoreboard
	cp ./bin/scoreboard /usr/local/bin/scoreboard

migration :
	sh ./scripts/migrations.sh

uninstall :
	rm /usr/local/bin/scoreboard
	rm ~/scoreboard.db
