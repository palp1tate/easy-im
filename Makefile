.PHONY: help
help:
	@echo "Usage:"
	@echo "  make help                   Show this help message"
	@echo "  make dev-up                 Start the development environment"
	@echo "  make dev-stop               Stop the development environment"
	@echo "  make dev-down               Clean the development environment"
	@echo "  make prod-up                Start the production environment"
	@echo "  make prod-stop              Stop the production environment"
	@echo "  make prod-down              Clean the production environment"
	@echo "  make etcd                   Get etcd key"
	@echo "  make status                 Show the status of the containers"


# Start development environment
.PHONY: dev-up
dev-up:
	docker-compose up -d

# Stop development environment
.PHONY: dev-stop
dev-stop:
	docker-compose stop

# Clean development environment
.PHONY: dev-down
dev-down:
	docker-compose down

# Start production environment
.PHONY: prod-up
prod-up:
	docker-compose -f docker-compose.prod.yml up -d

# Stop production environment
.PHONY: prod-stop
prod-stop:
	docker-compose -f docker-compose.prod.yml stop

# Clean production environment
.PHONY: prod-down
prod-down:
	docker-compose -f docker-compose.prod.yml down

# Get etcd key
.PHONY: etcd
etcd:
	docker exec -it easy-im-etcd etcdctl get --prefix ""

# Show the status of the containers
.PHONY: status
status:
	docker-compose ps