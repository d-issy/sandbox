import unittest

from pydantic import BaseModel, ValidationError, constr, conint
from pydantic import PositiveInt, NegativeInt


class Model(BaseModel):
    upper_str: constr(to_upper=True) = None  # type: ignore
    lower_str: constr(to_lower=True) = None  # type: ignore
    short_str: constr(min_length=2, max_length=10) = None  # type: ignore
    regex_str: constr(regex=r"^apple (pie|tart|sandwitch)$") = None  # type: ignore
    strip_str: constr(strip_whitespace=True) = None  # type: ignore

    mod_int: conint(multiple_of=5) = None  # type: ignore
    pos_int: PositiveInt = None  # type: ignore
    neg_int: NegativeInt = None  # type: ignore


class TestConstraintFieldTypes(unittest.TestCase):
    def test_upper_str(self):
        self.assertEqual(Model(upper_str="hello").upper_str, "HELLO")  # type: ignore

    def test_lower_str(self):
        self.assertEqual(Model(lower_str="HELLO").lower_str, "hello")  # type: ignore

    def test_shot_str(self):
        Model(short_str="ABCDEFG")

        with self.assertRaises(ValidationError):
            Model(short_str="ABCDEFGHIJKLMNOPQRSTUVWXYZ")

    def test_regex_str(self):
        Model(regex_str="apple pie")

        with self.assertRaises(ValidationError):
            Model(regex_str="ABCDEFGHIJKLMNOPQRSTUVWXYZ")

    def test_strip_str(self):
        self.assertEqual(Model(strip_str=" hello ").strip_str, "hello")  # type: ignore

    def test_mod_int(self):
        Model(mod_int=555)

        with self.assertRaises(ValidationError):
            Model(mod_int=556)

    def test_pos_int(self):
        Model(pos_int=1)

        with self.assertRaises(ValidationError):
            Model(pos_int=-1)

    def test_neg_int(self):
        Model(neg_int=-1)

        with self.assertRaises(ValidationError):
            Model(neg_int=1)
