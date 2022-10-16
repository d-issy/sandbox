RUNNER_PREFIX := docker-compose run --rm python

.PHONY: fmt build test
fmt:
	$(RUNNER_PREFIX) black .

build:
	docker-compose build

test:
	$(RUNNER_PREFIX) mypy . \
		--enable-incomplete-feature=TypeVarTuple \

	$(RUNNER_PREFIX) python -m unittest -v
