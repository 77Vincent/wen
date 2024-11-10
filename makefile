run:
	hugo serve

deploy:
	hugo deploy --target production && hugo deploy --target production-www

.PHONY: run, deploy
