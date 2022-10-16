import unittest

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


class TestValidator(unittest.TestCase):
    def test_ok(self):
        user = UserModel(
            name="d issy",
            username="issy",
        )
        self.assertEqual(user.name, "D Issy")
        self.assertEqual(user.username, "issy")

    def test_ng(self):
        with self.assertRaises(ValidationError):
            try:
                UserModel(
                    name="issy",
                    username="issy dayo!",
                )
            except ValidationError as e:
                self.assertEqual(e.errors()[0]["loc"], ("name",))
                self.assertEqual(e.errors()[0]["msg"], "must contain a space")
                self.assertEqual(e.errors()[0]["type"], "value_error")

                self.assertEqual(e.errors()[1]["loc"], ("username",))
                self.assertEqual(e.errors()[1]["msg"], "must be alphanumeric")
                self.assertEqual(e.errors()[1]["type"], "assertion_error")
                raise e
