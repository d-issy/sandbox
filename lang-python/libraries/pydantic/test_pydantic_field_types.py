import unittest

from pydantic import BaseModel, ValidationError

from pydantic import HttpUrl, FileUrl
from pydantic.color import Color
from pydantic import SecretStr, SecretBytes
from pydantic.types import PaymentCardBrand, PaymentCardNumber


class Model(BaseModel):
    http_url: HttpUrl = None  # type: ignore
    file_url: FileUrl = None  # type: ignore

    color: Color = None  # type: ignore

    password: SecretStr = None  # type: ignore
    password_bytes: SecretBytes = None  # type: ignore

    creditcard_number: PaymentCardNumber = None  # type: ignore


class TestPydanticFieldTypes(unittest.TestCase):
    def test_http_url(self):
        self.assertEqual(
            Model(http_url="https://github.com/d-issy/").http_url,  # type: ignore
            "https://github.com/d-issy/",
        )

    def test_http_ur_ng(self):
        with self.assertRaises(ValidationError):
            Model(http_url="invalid url").http_url  # type: ignore

    def test_file_http(self):
        self.assertEqual(
            Model(file_url="file:///root/hello.txt").file_url,  # type: ignore
            "file:///root/hello.txt",
        )

    def test_file_http_ng(self):
        with self.assertRaises(ValidationError):
            Model(file_url="https://github.com/d-issy/").file_url  # type: ignore

    def test_color(self):
        self.assertEqual(Model(color="#123456").color.as_hex(), "#123456")  # type: ignore
        self.assertEqual(Model(color="green").color.as_named(), "green")  # type: ignore

    def test_secret(self):
        self.assertEqual(Model(password="hello").password.get_secret_value(), "hello")  # type: ignore
        self.assertEqual(Model(password_bytes="hello").password_bytes.get_secret_value(), b"hello")  # type: ignore

    def test_creditcard_number(self):
        dummy_number = 4000000000000002
        self.assertEqual(Model(creditcard_number=dummy_number).creditcard_number.masked, "400000******0002")  # type: ignore
        self.assertEqual(Model(creditcard_number=dummy_number).creditcard_number.brand, PaymentCardBrand.visa)  # type: ignore
