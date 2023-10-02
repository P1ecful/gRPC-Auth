# Training grpc project

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

For starting you need to install protobufs. 
Instruction for windows: https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/

### Installation

Below is an example of how you can install a project
1. Get grpc
   ```sh
   $ export GO111MODULE=on
   $ go get google.golang.org/protobuf/cmd/protoc-gen-go
   $ go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
   ```
2. Clone the repo
   ```sh
   $ git clone https://github.com/P1ecful/Authorization-Service
   $
   ```
3. Go.mod
   ```sh
   $ go mod tidy
   $ 
   ```
 4. Create database with name "usersmeta.db". And create the table:
       ```sh
      $ CREATE TABLE "Meta" (
	   $   "userHash"	TEXT NOT NULL UNIQUE
      $ );
      $
     ```
 5. Run server
     ```sh
     $ go run cmd/server/main.go
     $
     ```

6. Go to postman, then select grpc and import protofile. Select the register method and submit your data. Example:
   ```sh
   $ {
   $    "login": "dolor",
   $    "password": "est aliquip Ut fugiat officia"
   $ }
   ```
