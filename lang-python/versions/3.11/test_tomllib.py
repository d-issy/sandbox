import tomllib


def test_toml():
    toml = """\
[tool.poetry]
name = "python311-test"
version = "0.1.0"
description = ""
authors = ["issy"]

[tool.poetry.dependencies]
python = "^3.11"
"""

    data = tomllib.loads(toml)
    assert data["tool"]["poetry"]["name"] == "python311-test"
    assert data["tool"]["poetry"]["version"] == "0.1.0"
    assert data["tool"]["poetry"]["description"] == ""
    assert len(data["tool"]["poetry"]["authors"]) == 1

    assert data["tool"]["poetry"]["dependencies"]["python"] == "^3.11"
