import unittest


class TestExceptionGroup(unittest.TestCase):
    def test_group(self):
        try:
            raise ExceptionGroup("examples", [ValueError(), TypeError()])
        except ExceptionGroup as eg:
            assert type(eg.exceptions[0]) is ValueError
            assert type(eg.exceptions[1]) is TypeError

    def test_group2(self):
        try:
            raise ExceptionGroup("examples", [ValueError(), TypeError()])
        except* ValueError:
            print("ValueError")
        except* TypeError:
            print("TypeError")
