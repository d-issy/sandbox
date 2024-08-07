from typing import Deque, Sequence

from pydantic import BaseModel


class Model(BaseModel):
    simple_list: list = None  # type: ignore
    list_of_ints: list[int] = None  # type: ignore

    simple_tuple: tuple = None  # type: ignore
    tuple_of_different_types: tuple[int, float, str, bool] = None  # type: ignore

    simple_dict: dict = None  # type: ignore
    dict_str_float: dict[str, float] = None  # type: ignore

    simple_set: set = None  # type: ignore
    set_bytes: set[bytes] = None  # type: ignore
    frozen_set: frozenset[int] = None  # type: ignore

    str_or_bytes: str | bytes = None  # type: ignore
    none_or_str: str | None = None  # type: ignore

    sequence_of_ints: Sequence[int] = None  # type: ignore

    compound: dict[str | bytes, list[set[int]]] = None  # type: ignore

    deque: Deque[int] = None  # type: ignore


def test_simple_list():
    assert Model(simple_list=["1", "2", "3"]).simple_list == ["1", "2", "3"]


def test_list_of_ints():
    assert Model(list_of_ints=["1", "2", "3"]).list_of_ints == [1, 2, 3]  # type:ignore


def test_simple_tuple():
    assert Model(simple_tuple=("1", "2", "3")).simple_tuple == ("1", "2", "3")  # type: ignore


def test_tuple_of_different_types():
    assert Model(tuple_of_different_types=("1", "2.3", "str", "true")).tuple_of_different_types == (1, 2.3, "str", True)  # type: ignore


def test_simple_dict():
    assert Model(simple_dict={"a": 1, "b": 2}).simple_dict == {"a": 1, "b": 2}


def test_str_float():
    assert Model(dict_str_float={"a": 1, "b": 2}).dict_str_float == {"a": 1.0, "b": 2.0}


def test_simple_set():
    assert Model(simple_set={"1", "2"}).simple_set == {"1", "2"}  # type: ignore
