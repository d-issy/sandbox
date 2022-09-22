from typing import Self

# mypyではまだ対応されていない at 2022.09
# https://peps.python.org/pep-0673/
# ref: https://github.com/python/mypy/pull/11666

# class Clonable:
#     def __init__(self, value: str) -> None:
#         self.value = value
#
#     def set_value(self, value: str) -> Self:
#         self.value = value
#         return self
