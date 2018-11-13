build:
	docker build -t messari/redirect-to-https .

push:
	docker push messari/redirect-to-https:latest
