RUNNER_PREFIX := docker-compose run --rm python

.PHONY: fmt
fmt:
	$(RUNNER_PREFIX) black .

.PHONY: build
build:
	docker-compose build

.PHONY: test
test:
	$(RUNNER_PREFIX) mypy .
	$(RUNNER_PREFIX) python -m unittest -v

.PHONY: shell
shell:
	$(RUNNER_PREFIX) ash

