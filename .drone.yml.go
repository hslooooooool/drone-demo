workspace:
  base: /srv/drone-demo
  path: .

pipeline:
  build:
     image: golang:alpine
     # pull: true
     environment:
       - KEY=VALUE
     secrets: [key1, key2]
     commands:
       - echo $$KEY
       - pwd
       - ls
       - CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
       - ./app