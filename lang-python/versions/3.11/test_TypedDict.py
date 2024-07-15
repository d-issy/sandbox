import unittest
from typing import NotRequired, Required, TypedDict

# see: https://peps.python.org/pep-0655/


class TodoA(TypedDict):
    title: str
    completed: NotRequired[bool]


class TodoB(TypedDict, total=False):
    title: Required[str]
    completed: bool


class TestTypedDict(unittest.TestCase):
    def test_not_require_ok(self):
        ok_1: TodoA = {"title": "TypedDictを試してみる", "completed": True}
        ok_2: TodoA = {"title": "TypedDictを試してみる"}

    def test_not_require_ng(self):
        # ng: TodoA = {}
        ...

    def test_require_ok(self):
        ok_1: TodoB = {"title": "TypedDictを試してみる", "completed": True}
        ok_2: TodoB = {"title": "TypedDictを試してみる"}

    def test_require_ng(self):
        # ng: TodoB = {}
        ...
