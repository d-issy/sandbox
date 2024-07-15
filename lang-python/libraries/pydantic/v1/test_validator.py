import pytest
from pydantic import BaseModel, ValidationError, validator


class UserModel(BaseModel):
    name: str
    username: str

    @validator("name")
    def name_must_contains_space(cls, v):
        if " " not in v:
            raise ValueError("must contain a space")
        return v.title()

    @validator("username")
    def username_alphanumeric(cls, v):
        assert v.isalnum(), "must be alphanumeric"
        return v


class TestValidator:
    def test_ok(self):
        user = UserModel(
            name="d issy",
            username="issy",
        )
        assert user.name == "D Issy"
        assert user.username == "issy"

    def test_ng(self):
        with pytest.raises(ValidationError):
            try:
                UserModel(
                    name="issy",
                    username="issy dayo!",
                )
            except ValidationError as e:
                assert e.errors()[0]["loc"] == ("name",)
                assert e.errors()[0]["msg"] == "must contain a space"
                assert e.errors()[0]["type"] == "value_error"

                assert e.errors()[1]["loc"] == ("username",)
                assert "must be alphanumeric" in e.errors()[1]["msg"]
                assert e.errors()[1]["type"] == "assertion_error"
                raise e
