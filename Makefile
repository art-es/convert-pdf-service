GOIMAGE=svs-go-app
NODEIMAGE=svs-node-app
DNETWORK=script-vs-server

build:
	docker build \
		-f ./docker/go.Dockerfile \
		-t ${GOIMAGE} \
		.
	
	docker build \
		-f ./docker/node.Dockerfile \
		-t ${NODEIMAGE} \
		.


run-go-app:
	docker run \
		--rm \
		-v $(CURDIR)/${APPDIR}:/app \
		-v $(CURDIR)/html:/html \
		-v $(CURDIR)/pdf:/pdf \
		--network ${DNETWORK} \
		${GOIMAGE}
# rm ./pdf/base.html


network:
	docker network create ${DNETWORK}


run-script:
	make run-go-app APPDIR=script


run-server:
	make run-go-app APPDIR=server


# docker run -v=(pwd)/converter:/converter -p 8080:8080 --network script-vs-server --network-alias=convert -it svs-node-app sh