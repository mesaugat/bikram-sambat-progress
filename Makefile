build:
	docker build -t bikram-sambat-progress .

run:
	docker run \
		-e TWITTER_CONSUMER_KEY=${TWITTER_CONSUMER_KEY} \
		-e TWITTER_CONSUMER_SECRET=${TWITTER_CONSUMER_SECRET} \
		-e TWITTER_ACCESS_TOKEN=${TWITTER_ACCESS_TOKEN} \
		-e TWITTER_ACCESS_TOKEN_SECRET=${TWITTER_ACCESS_TOKEN_SECRET} \
		bikram-sambat-progress

rund:
	docker run -dit \
		--restart unless-stopped \
		-e TWITTER_CONSUMER_KEY=${TWITTER_CONSUMER_KEY} \
		-e TWITTER_CONSUMER_SECRET=${TWITTER_CONSUMER_SECRET} \
		-e TWITTER_ACCESS_TOKEN=${TWITTER_ACCESS_TOKEN} \
		-e TWITTER_ACCESS_TOKEN_SECRET=${TWITTER_ACCESS_TOKEN_SECRET} \
		bikram-sambat-progress

test:
	go test -v ./...

.DEFAULT_GOAL := build
