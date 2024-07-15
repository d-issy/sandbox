import pytest
from pydantic import (
    BaseModel,
    FileUrl,
    HttpUrl,
    SecretBytes,
    SecretStr,
    ValidationError,
)
from pydantic.color import Color
from pydantic.types import PaymentCardBrand, PaymentCardNumber


class Model(BaseModel):
    http_url: HttpUrl = None  # type: ignore
    file_url: FileUrl = None  # type: ignore

    color: Color = None  # type: ignore

    password: SecretStr = None  # type: ignore
    password_bytes: SecretBytes = None  # type: ignore

    creditcard_number: PaymentCardNumber = None  # type: ignore


def test_http_url():
    assert Model(http_url="https://github.com/d-issy/").http_url == "https://github.com/d-issy/"  # type: ignore


def test_http_ur_ng():
    with pytest.raises(ValidationError):
        Model(http_url="invalid url").http_url  # type: ignore


def test_file_http():
    assert Model(file_url="file:///hello.txt").file_url == "file:///hello.txt"  # type: ignore


def test_file_http_ng():
    with pytest.raises(ValidationError):
        Model(file_url="https://github.com/d-issy/").file_url  # type: ignore


def test_color():
    assert Model(color="#123456").color.as_hex() == "#123456"  # type: ignore
    assert Model(color="green").color.as_named() == "green"  # type: ignore


def test_secret():
    assert Model(password="hello").password.get_secret_value() == "hello"  # type: ignore
    assert Model(password_bytes="hello").password_bytes.get_secret_value() == b"hello"  # type: ignore


def test_creditcard_number():
    dummy_number = 4000000000000002
    assert Model(creditcard_number=dummy_number).creditcard_number.masked == "400000******0002"  # type: ignore
    assert Model(creditcard_number=dummy_number).creditcard_number.brand == PaymentCardBrand.visa  # type: ignore
