# Task 1

## Docker

The run.sh script handles the setup, build, and running of the application in a Docker container. This is the
recommended approach for easy setup.

```shell
./run.sh
```

## go build

```shell
go run cmd/api/main.go
```

## API Endpoints

The frontend JavaScript communicates with two main backend endpoints for key generation and extraction. Each endpoint is
designed to accept JSON input and return JSON responses.

#### 1. Extract Public Key Endpoint

Endpoint: POST /v1/extract-public-key
Description: Accepts a private key and returns the associated public key if the private key is valid.

##### Request Format:

POST /v1/extract-public-key
Content-Type: application/json

Request Body:

```json

{
"private_key": "<YourPrivateKey>"
}
```
##### Expected Response:

On Success:

```json

{
  "public_key": "<YourPublicKey>"
}
```

#### 2. Generate Private Key Endpoint

Endpoint: POST /v1/generate-private-key
Description: Generates a new private key based on provided parameters like name, email, and bit length.

##### Request Format:

POST /v1/generate-private-key
Content-Type: application/json

Request Body:

```json

{
  "name": "Test User",
  "email": "test@example.com",
  "bit_length": 2048
}
```

##### Expected Response:

On Success:

```json

{
  "private_key": "<GeneratedPrivateKey>"
}
``` 

## Usage

Please open http://localhost:3002/