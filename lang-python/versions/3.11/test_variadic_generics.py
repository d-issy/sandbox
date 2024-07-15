# ref: https://peps.python.org/pep-0646/

import unittest
from typing import Generic, TypeVar, TypeVarTuple

R = TypeVar("R")
G = TypeVar("G")
B = TypeVar("B")


class Color(Generic[R, G, B]):
    def __init__(self, R, G, B) -> None:
        super().__init__()
        self.R = R
        self.G = G
        self.B = B


# TypeVarTuple は mypy で対応されていない at 2022.09
# ref: https://github.com/python/mypy/issues/12280

Shape = TypeVarTuple("Shape")
class Array(Generic[*Shape]):
    def __init__(self, shape: tuple[*Shape]) -> None:
        self.shape: tuple[*Shape] = shape


class TestVariadicGenerics(unittest.TestCase):
    def test_TypeVar(self):
        _: Color[int, int, int] = Color(0, 0, 0)

    def test_TypeVarTuple(self):
        _: Array[int, int] = Array((0, 0))
