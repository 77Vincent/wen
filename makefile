run:
	hugo serve

deploy:
	hugo build && hugo deploy --target production && hugo deploy --target production-www

index:
	npm run index-and-send

.PHONY: run, deploy
