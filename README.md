**A User Management System** written in Golang with gRPC protocol 
to match CRUD. There are also an admin and a moderator users that can manage user`s data.


# Tech-Stack
* go 1.20
* PostgreSQL
* [env](https://github.com/caarlos0/env)
* [gRPC](github.com/golang/protobuf)
* [gRPC-validation](github.com/mwitkow/go-proto-validators)
* [logger](https://github.com/rs/zerolog)


# Getting Started
```
cd existing_repo
git remote add origin https://github.com/BloodsFa1zer/grpc-crud
git branch -M main
git push -uf origin main

cd ./proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --govalidators_out=paths=source_relative:. user.proto
```

