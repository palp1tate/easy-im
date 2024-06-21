.PHONY: help
help:
	@echo "Usage:"
	@echo "  make help                   Show this help message"
	@echo "  make dev-up                 Start the development environment"
	@echo "  make dev-stop               Stop the development environment"
	@echo "  make dev-down               Clean the development environment"
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

# Get etcd key
.PHONY: etcd
etcd:
	docker exec -it easy-im-etcd etcdctl get --prefix ""

# Show the status of the containers
.PHONY: status
status:
	docker-compose ps