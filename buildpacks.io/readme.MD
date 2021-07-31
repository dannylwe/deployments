# Creating an image using buildpacks
## To view options
 ```
$ pack
```

## To see available builders
```
$ pack builder suggest
```

## Build image from source using builder
### Example:
- pack build {image-name} --builder {builder}  
```
$ pack build hello-golang --builder paketobuildpacks/builder:tiny
```

## Start and check application
```
$ docker run --rm -p 9001:9001 hello-golang
```
- confirm application is running `curl localhost:9001/health`