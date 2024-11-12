run:
	hugo serve

deploy:
	hugo deploy --target production && hugo deploy --target production-www

index:
	npm run index-and-send

.PHONY: run, deploy
