# scale時のpod数を指定する
NUM = 1

.PHONY: build apply delete clean
build:
	docker build -f locst.docker -t locust .

apply: build
	kubectl apply -f manifests/web.yaml
	kubectl apply -f manifests/locust.yaml
	@echo
	@echo Locust Web UI URL: http://localhost/

delete:
	kubectl delete -f manifests/web.yaml
	kubectl delete -f manifests/locust.yaml

clean:
	-docker rmi locust
	-docker image prune -f

.PHONY: watch
watch:
	watch kubectl get deploy,svc,rs -o wide

.PHONY: scale/*
scale/web:
	kubectl scale --replicas $(NUM) deploy/web

scale/worker:
	kubectl scale --replicas $(NUM) deploy/locust-worker
