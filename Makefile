test:
	go test .

build: test
	go build -o cloudbuild-slack-notifier .

run: build
	chmod +x cloudbuild-slack-notifier
	./cloudbuild-slack-notifier

build-docker: test
	docker build . --tag cloudbuild-slack-notifier:latest

deploy: test
	./deployment.sh
