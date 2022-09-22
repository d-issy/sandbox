from typing import Iterable
import unittest
from typing import LiteralString

# ref: https://peps.python.org/pep-0675/


class Connection:
    def execute(self, query: LiteralString, parameters: Iterable[str]):
        f"execute query: `{query}` with {parameters}"


class TestArbitraryLiteralStringType(unittest.TestCase):
    def test_ok(self):
        conn = Connection()
        query = "SELECT * FROM users WHERE user_id = ?"
        conn.execute(query, ("issy",))

    # https://github.com/python/mypy/issues/12554
    @unittest.skip("Not yet supported by mypy. so Passed now")
    def test_ng(self):
        conn = Connection()
        user_id = "issy"
        query = f"SELECT * FROM users WHERE user_id = {user_id}"
        conn.execute(query, ())
