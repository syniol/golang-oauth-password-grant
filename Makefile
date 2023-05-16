help:
	@awk 'BEGIN {FS = ":.*#"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?#/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^#@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


.PHONY: deploy
deploy: # Deploys Services for OAuth Authentication
	docker-compose -f ./deploy/docker/docker-compose.yml build
	docker-compose -f ./deploy/docker/docker-compose.yml up -d


.PHONY: destroy
destroy: # Destroys Services for OAuth Authentication
	docker-compose -f ./deploy/docker/docker-compose.yml down


.PHONY: nuke
nuke: # Nukes all docker facilities for OAuth Authentication
	docker-compose -f ./deploy/docker/docker-compose.yml down --rmi local -v


.PHONY: debug # Runs in debug mode
debug:
	docker-compose -f ./deploy/docker/docker-compose.yml -f ./deploy/docker/docker-compose.debug.yml up -d database cache