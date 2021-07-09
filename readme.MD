## Endpoints
`curl localhost:9001/health`  
`curl localhost:9001/hello`

## build with docker
docker build -t hello .  

## run with docker
docker run --rm -p 9001:9001 hello

## Deploy with heroku 
- Login into Heroku with `heroku login`  
- Login into Heroku registry `heroku container:login`    
- Create a Heroku app `heroku create`  
- Push to registry `heroku container:push web`  
- Push release `heroku container:release web`
