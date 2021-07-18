## Prerequistes
- docker
- aws account
- IAM user with administrator access
- latest AWS CLI configured

Note: To use AWS ECR use [ecr_aws](https://docs.aws.amazon.com/AmazonECR/latest/userguide/getting-started-cli.html)

## Authenticate ECR registry
- `aws ecr get-login-password --region {{region}} | docker login --username AWS --password-stdin {{aws_account_id}}.dkr.ecr.{{region}}.amazonaws.com`   

- example: 
```
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 269757654257.dkr.ecr.us-east-1.amazonaws.com
````

- Go to Amazon ECR on the browser and create a repository or run  
```
aws ecr create-repository \
    --repository-name hello-world \
    --image-scanning-configuration scanOnPush=true \
    --region us-east-1
```
where "hello-world" is the name of the repository

## Push Image to repository  
- Tag image
```
docker tag hello:v1 {{aws_account_id}}.dkr.ecr.us-east-1.amazonaws.com/hello-world:v1
```
- Push image
```
docker push {{aws_account_id}}.dkr.ecr.us-east-1.amazonaws.com/hello-world:v1
```

## (Optional) 
- Pull image from ECR
```
docker pull {{aws_account_id}}.dkr.ecr.us-east-1.amazonaws.com/hello-world:v1
```
- Delete Image
```
aws ecr batch-delete-image \
      --repository-name hello-world \
      --image-ids imageTag=v1
```
- Delete Repo
```
aws ecr delete-repository \
      --repository-name hello-world \
      --force
```
Note: Take note of the image URI in the repo  
269757654257.dkr.ecr.us-east-1.amazonaws.com/hello-world:v1

## create cluster
## create Task using CLI
- [Task](https://docs.amazonaws.cn/en_us/AmazonECS/latest/developerguide/ECS_AWSCLI_EC2.html)

## create a task (follow the prompt information)
- After the task(s) has been deployed. Go to ec2 (if used) and on the security group, open the PORTS used by the application.
- use the public IP of the ec2 machine on the open PORT to see/use the application 

## Four Steps to Run ECS Clusters on EC2 Spot Instances
[spot-instances](https://aws.amazon.com/ec2/spot/containers-for-less/get-started/)
