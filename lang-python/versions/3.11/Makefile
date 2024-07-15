setup:
	poetry install

test:
	poetry run pytest
	poetry run mypy . --check-untyped-defs


format:
	poetry run black .
