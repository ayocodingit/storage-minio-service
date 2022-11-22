<a href="https://codeclimate.com/github/ayocodingit/storage-minio-service/maintainability"><img src="https://api.codeclimate.com/v1/badges/9b4a34093b55bc0e2823/maintainability" /></a>
[![GitHub issues](https://img.shields.io/github/issues/ayocodingit/storage-minio-service)](https://github.com/ayocodingit/storage-minio-service/issues)
[![GitHub forks](https://img.shields.io/github/forks/ayocodingit/storage-minio-service)](https://github.com/ayocodingit/storage-minio-service/network)
[![GitHub stars](https://img.shields.io/github/stars/ayocodingit/storage-minio-service)](https://github.com/ayocodingit/storage-minio-service/stargazers)


# Storage Minio Service

## Detail information for endpoints 

There are several endpoints that can be used, including:

```curl
$ POST /upload              * endpoint for upload file
$ GET /download/:filename   * endpoint for get file (if configuration cloud set private)
$ DELETE /delete/:filename  * endpoint for remove file
```

## Tech Stacks

- **Golang v1.17** - <https://go.dev/>
- **Gin** - <https://gin-gonic.com/>
- **Minio** - <https://min.io/>

## Quick Start

Clone the project:

```bash
$ git clone https://github.com/ayocodingit/storage-minio-service.git
$ cd storage-minio-service
$ cp .env.example .env
```

## How to Run on local
```bash
$ go run src/cmd/main.go
```

## How to Build docker
```bash
$ docker build -f docker/Dockerfile -t storage-minio-service:<VERSION> .
```
