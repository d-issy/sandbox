# load-test locust

## requirements
- docker
- kubernetes

## how to play

```
make apply
```

and then, access http://localhost/



## cleanup

```
make delete
make clean
```


## scaling

```
make scale/worker NUM={NUMER}
make scale/web NUM={NUMER}
```

