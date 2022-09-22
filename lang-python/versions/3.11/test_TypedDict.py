import unittest
from typing import TypedDict, Required, NotRequired

# ref: https://peps.python.org/pep-0655/


class TodoA(TypedDict):
    title: str
    completed: NotRequired[bool]


class TodoB(TypedDict, total=False):
    title: Required[str]
    completed: bool


class TestTypedDict(unittest.TestCase):
    def test_not_require_ok(self):
        _: TodoA = {"title": "TypedDictを試してみる", "completed": True}
        _: TodoA = {"title": "TypedDictを試してみる"}

    def test_not_require_ng(self):
        pass
        # mypyエラー: error: Missing key "title" for TypedDict "TodoA"
        # _: TodoA = {}

    def test_require_ok(self):
        _: TodoB = {"title": "TypedDictを試してみる", "completed": True}
        _: TodoB = {"title": "TypedDictを試してみる"}

    def test_require_ng(self):
        pass
        # mypyエラー: error: Missing key "title" for TypedDict "TodoB"
        # _: TodoB = {}
