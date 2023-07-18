# Usage

## Method #1 (Makefile)

```sh
make run_client
make run_server
```

## Method #2 (docker-compose)

```sh
docker-compose up -d

docker logs pow-wow_client_1 -f
# or
docker logs pow-wow_server_1 -f
```

## Environment variables

### Client

| name           | type    | default        | description
|----------------|---------|----------------|--------------------------------------
| SERVER_ADDR    | string  | 127.0.0.1:9000 | listen TCP address
| FETCH_WORKERS  | int     | 4              | count of client requests at same time
| TIMEOUT        | int     | 1000           | timeout after failed request

### Server

| name             | type    | default        | description
|------------------|---------|----------------|----------------------------------------
| LISTEN_ADDR      | string  | 0.0.0.0:9000   | server TCP address
| DIFFICULTY       | byte    | 23             | difficulty of calc algorithm for client
| PROOF_TOKEN_SIZE | int     | 64             | data size for proof calc for client

# Implementation

## Project structure

- deploy - dockerfiles for build
- cmd
  - server - server side app
  - client - client side app
- internal
  - pow - implementation "Proof Of Work" algorithm based on sha256
  - client - implementation "Proof Of Work" client requests
  - server - implementation "Proof Of Work" server listener

## Challenge-Response protocol

- Client connected to server
- Server write connection to log
- Server send puzzle packet
  | offset             | name         | length
  | -------------------|--------------|---------------
  |                  0 | difficulty   | 1 byte
  |                  1 | token size   | 2 bytes
  |                  3 | rand token   | ProofTokenSize
- Client calculate proof based on difficulty and data (rand token)
- Client send proof nonce to server
- Server check proof based on difficulty, data (rand token) and received nonce
- If proof is valid
  - write to log
  - server send to client quote from “Word of Wisdom”
  - client print response from server
  - connection close
- If proof is not valid
  - write log log
  - connection close

## Algorithm

I chose the sha256 algorithm because:
- It presents in the standard go library
- "nonce" number not too big
- Easy calculate zeroes in hash
- The difficulty of the calculation is enough to guard from ddos

### Calculation hash basis

| offset | name              | length
|--------|-------------------|----------------------------------
| 0      | nonce             | 8 bytes
| 8      | data (rand token) | ProofTokenSize (usually 64 bytes)
