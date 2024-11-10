run:
	hugo serve

deploy:
	hugo deploy --target production --target production-www

.PHONY: run, deploy
