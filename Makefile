build:
	docker build -t bikram-sambat-progress .

run:
	docker run bikram-sambat-progress

rund:
	docker run -d bikram-sambat-progress

.DEFAULT_GOAL := build
