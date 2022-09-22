RUNNER_PREFIX := docker-compose run --rm python

.PHONY: fmt test
fmt:
	$(RUNNER_PREFIX) black .

test:
	$(RUNNER_PREFIX) mypy .
	$(RUNNER_PREFIX) python -m unittest -v
