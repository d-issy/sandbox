import unittest
from datetime import datetime

from pydantic import BaseModel, ValidationError


class User(BaseModel):
    id: int
    name: str
    signup_ts: datetime | None = None
    friends: list[int] = []


class TestModel(unittest.TestCase):
    def test_ok(self):
        data = {
            "id": "123",
            "name": "issy",
            "friends": [1, 2, "3"],
        }
        user = User(**data)
        self.assertEqual(user.id, 123)
        self.assertEqual(user.friends[2], 3)

    def test_validation(self):
        with self.assertRaises(ValidationError):
            try:
                User(signup_ts="broken", friends=[1, 2, "not number"])  # type: ignore
            except ValidationError as e:
                self.assertEqual(e.errors()[0]["loc"], ("id",))
                self.assertEqual(e.errors()[0]["msg"], "field required")
                self.assertEqual(e.errors()[0]["type"], "value_error.missing")

                self.assertEqual(e.errors()[1]["loc"], ("name",))
                self.assertEqual(e.errors()[1]["msg"], "field required")
                self.assertEqual(e.errors()[1]["type"], "value_error.missing")

                self.assertEqual(e.errors()[2]["loc"], ("signup_ts",))
                self.assertEqual(e.errors()[2]["msg"], "invalid datetime format")
                self.assertEqual(e.errors()[2]["type"], "value_error.datetime")

                self.assertEqual(e.errors()[3]["loc"], ("friends", 2))
                self.assertEqual(e.errors()[3]["msg"], "value is not a valid integer")
                self.assertEqual(e.errors()[3]["type"], "type_error.integer")
                raise e
