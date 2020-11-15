# Description
Program will run new ec2 instances, wait for status checks will pass and kill instances with an old image.

# Build
```shell script
go build
```

# Run
```shell script
./deploy {old ami} {new ami}
```
