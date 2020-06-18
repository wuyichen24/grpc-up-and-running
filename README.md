# grpc-up-and-running

[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](https://opensource.org/licenses/Apache-2.0) 

The study note of the book "[gRPC: Up and Running (Kasun Indrasiri)](http://shop.oreilly.com/product/0636920282754.do)" and the reconstruction of source code.

![](imgs/book-cover.jpg)

## Study Notes
- [Google Doc](https://docs.google.com/document/d/1-9X1T80fF26CSScx9xJNqLx5o1ruhYRGycqEnd9W5Gk/edit?usp=sharing)

## Experience 
- [**Install Protocol Buffer Compiler**](docs/install_protocol_buffer_compiler.md)
- [**Generate Server Stub**](docs/generate_stub_go.md)
- [**Build Executable File**](docs/build_executable.md)
- [**Write Client Code**](docs/write_client.md)
   - [Implement Main Method](docs/write_client.md#implement-main-method)
- [**Write Server Code**](docs/write_server.md)
   - [Implement Main Method](docs/write_server.md#implement-main-method)
   - [Implement Remote Methods](docs/write_server.md#implement-remote-methods)
- [**Authentication**](docs/authentication.md)
   - [TLS Authentication](docs/authentication.md#tls-authentication)
      - [One-way TLS](docs/authentication.md#one-way-tls)
      - [Two-way TLS (mTLS)](#two-way-tls-mtls)
   - [Other Authentication Solutions](docs/authentication.md#other-authentication-solutions)
      - [Basic Authentication](docs/authentication.md#basic-authentication)
      - [OAuth 2.0](docs/authentication.md#oauth-20)
      - [JWT](docs/authentication.md#jwt)
      - [Google Token-Based Authentication](docs/authentication.md#google-token-based-authentication)
- [**gRPC Gateway**](docs/grpc_gateway.md)

## Directory Structure
- **docs**: The study notes of this books
   - diagram: The diagrams for this repository.
- **examples**: The example code of gRPC sub-techniques.
   - loadbalancing: The load balancer for multiple gRPC services.
   - **security**: The example of the authentication solutions for gRPC.
      - one-way-tls: The one-way TLS authentication.
      - two-way-tls: The two-way (mTLS) authentication.
      - basic-auth: The basic authentication.
      - oauth2: The OAuth 2 authentication.
      - jwt: The JWT authentication.
   - grpc-gateway: The gRPC gateway example.
- **imgs**: The images for this repository.
- **productinfo**: The hello-world example of gRPC.
- **ordermgt**: The gRPC examples for demostrating 4 gRPC communication patterns.

## Differences to The Original Source Code
- Add the detailed [instruction](docs/install_protocol_buffer_compiler.md) about how to install protocol buffer compiler.
- Add tutorials of writing server code and client code and modularize them by functionality.
- Flatten the source code by chapter into one application.
- Make sure the code is runnable (Fix some issues in the original source code).
- Better documentation.
- Add comments to make the code easy to read.

## Services And Remote Methods
### Product Info

![](docs/diagram/productinfo.png)

| Method | Pattern | Description | 
|---|---|---|
| AddProduct | Unary RPC | Add a product. |
| GetProduct | Unary RPC | Get a product by product ID. |

### Order Management

![](docs/diagram/ordermgt.png)

| Method | Pattern | Description | 
|---|---|---|
| AddOrder | Unary RPC | Add a new order. |
| GetOrder | Unary RPC | Get a order by order ID. |
| SearchOrders | Server-side streaming | Get all the orders which has a certain item. |
| UpdateOrders | Client-side streaming | Update multiple orders. |
| ProcessOrders | Bidirectional streaming | Process multiple orders. <li>All the order IDs will be sent from client as a stream.<li>A combined shipment will contains all the orders which will be delivered to the same destination.<li>When the max batch size is reached, all the currently created combined shipments will be sent back to the client. |
