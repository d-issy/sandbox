from typing import Iterable, LiteralString

# see: https://peps.python.org/pep-0675/


class Connection:
    def execute(self, query: LiteralString, parameters: Iterable[str]):
        f"execute query: `{query}` with {parameters}"


def test_ok():
    conn = Connection()
    query = "SELECT * FROM users WHERE user_id = ?"
    conn.execute(query, ("issy",))


def test_ng():
    conn = Connection()
    user_id = "issy"
    query = f"SELECT * FROM users WHERE user_id = {user_id}"
    conn.execute(query, ())
