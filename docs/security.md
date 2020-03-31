# Security

- [**TLS Authentication**]()
   - [**One-way TLS**]()
      - [**Server Code**]()
      - [**Client Code**]()
   - [**Two-way TLS**]()
      - [**Server Code**]()
      - [**Client Code**]()
- [**Other Authentication Solutions**]()
   - [**Basic Auth**]()
   - [**OAuth 2.0**]()
   - [**JWT**]()
   - [**Google Token-Based Authentication**]()

## TLS Authentication
### One-way TLS
#### Server Code
- Read and parse a public/private key pair and create a certificate.
  ```go
  cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
  ```
- Add an option to enable TLS for all incoming connections by adding certificates as TLS server credentials.
  ```
  opts := []grpc.ServerOption{
      grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
  }
  ```

#### Client Code


### Two-way TLS
#### Server Code
#### Client Code

## Other Authentication Solutions
### Basic Authentication
### OAuth 2.0
### JWT
### Google Token-Based Authentication
