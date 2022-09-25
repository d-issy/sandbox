import locust
from locust import HttpUser, task


class LoadTest(HttpUser):
    host = "http://web"

    @task
    def do(self):
        self.client.get('/')


