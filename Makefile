.PHONY: deploy rollback

deploy:
	docker compose --file ./deploy/docker/docker-compose.yml  up -d

rollback:
	docker compose --file ./deploy/docker/docker-compose.yml  down
	docker rmi docker-migration docker-avito-shop
