disable-xhost-security:
	xhost + # simplified solution, just for development

enable-xhost-security:
	xhost -

build-docker:
	docker build -t tic-cat-dog-client .

run-tests: disable-xhost-security build-docker
	docker run -e DISPLAY=${DISPLAY} -v /tmp/.X11-unix:/tmp/.X11-unix tic-cat-dog-client-build go test /app/src/...

run-game: build-docker
	docker run -e ASSETS_ABS_PATH=/app/assets -e DISPLAY=${DISPLAY} -v /tmp/.X11-unix:/tmp/.X11-unix tic-cat-dog-client /app/tic-cat-dog-client
