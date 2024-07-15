from datetime import datetime

import pytest
from pydantic import BaseModel, ValidationError


class User(BaseModel):
    id: int
    name: str
    signup_ts: datetime | None = None
    friends: list[int] = []


def test_ok():
    data = {
        "id": "123",
        "name": "issy",
        "friends": [1, 2, "3"],
    }
    user = User(**data)
    assert user.id == 123
    assert user.friends[2] == 3


def test_validation():
    with pytest.raises(ValidationError):
        try:
            User(signup_ts="broken", friends=[1, 2, "not number"])  # type: ignore
        except ValidationError as e:
            assert e.errors()[0]["loc"] == ("id",)
            assert e.errors()[0]["msg"] == "field required"
            assert e.errors()[0]["type"] == "value_error.missing"

            assert e.errors()[1]["loc"] == ("name",)
            assert e.errors()[1]["msg"] == "field required"
            assert e.errors()[1]["type"] == "value_error.missing"

            assert e.errors()[2]["loc"] == ("signup_ts",)
            assert e.errors()[2]["msg"] == "invalid datetime format"
            assert e.errors()[2]["type"] == "value_error.datetime"

            assert e.errors()[3]["loc"], "friends" == 2
            assert e.errors()[3]["msg"] == "value is not a valid integer"
            assert e.errors()[3]["type"] == "type_error.integer"
            raise e
