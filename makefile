run:
	hugo serve --cleanDestinationDir

deploy:
	rm -rf public && hugo build && hugo deploy

index:
	cd dev/indexNow && go run main.go

.PHONY: run, deploy, index
