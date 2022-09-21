test:
	go test .

build: test
	go build -o cloudbuild-slack-notifier .

run: build
	chmod +x cloudbuild-slack-notifier
	./cloudbuild-slack-notifier

deploy: test
	./deployment.sh
