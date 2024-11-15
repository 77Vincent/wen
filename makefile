run:
	hugo serve

deploy:
	rm -rf public && hugo build && hugo deploy

index:
	npm run index-and-send

.PHONY: run, deploy
