import pytest
from pydantic import (
    BaseModel,
    NegativeInt,
    PositiveInt,
    ValidationError,
    conint,
    constr,
)


class Model(BaseModel):
    upper_str: constr(to_upper=True) = None  # type: ignore
    lower_str: constr(to_lower=True) = None  # type: ignore
    short_str: constr(min_length=2, max_length=10) = None  # type: ignore
    regex_str: constr(regex=r"^apple (pie|tart|sandwitch)$") = None  # type: ignore
    strip_str: constr(strip_whitespace=True) = None  # type: ignore

    mod_int: conint(multiple_of=5) = None  # type: ignore
    pos_int: PositiveInt = None  # type: ignore
    neg_int: NegativeInt = None  # type: ignore


def test_upper_str():
    assert Model(upper_str="hello").upper_str == "HELLO"


def test_lower_str():
    assert Model(lower_str="HELLO").lower_str == "hello"


def test_shot_str():
    # no raises
    Model(short_str="ABCDEFG")

    with pytest.raises(ValidationError):
        Model(short_str="ABCDEFGHIJKLMNOPQRSTUVWXYZ")


def test_regex_str():
    # no raises
    Model(regex_str="apple pie")

    with pytest.raises(ValidationError):
        Model(regex_str="ABCDEFGHIJKLMNOPQRSTUVWXYZ")


def test_strip_str():
    assert Model(strip_str=" hello ").strip_str == "hello"


def test_mod_int():
    # no raises
    Model(mod_int=555)

    with pytest.raises(ValidationError):
        Model(mod_int=556)


def test_pos_int():
    # no raises
    Model(pos_int=1)

    with pytest.raises(ValidationError):
        Model(pos_int=-1)


def test_neg_int():
    # no raises
    Model(neg_int=-1)

    with pytest.raises(ValidationError):
        Model(neg_int=1)
