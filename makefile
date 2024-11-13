run:
	hugo serve

deploy:
	hugo build && hugo deploy

index:
	npm run index-and-send

.PHONY: run, deploy
