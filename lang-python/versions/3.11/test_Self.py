from typing import Self

# see: https://peps.python.org/pep-0673/


class Clonable:
    def __init__(self, value: str) -> None:
        self.value = value

    def set_value(self, value: str) -> Self:
        self.value = value
        return self
