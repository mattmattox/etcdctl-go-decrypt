# etcdctl-go-decrypt

This project is a Go-based tool designed to decrypt secrets stored in etcd. It includes AES encryption/decryption functionality and integrates with `etcdctl` for interacting with etcd.

## Features

- Decrypt AES-encrypted secrets from etcd.
- Static code analysis using `golint`, `staticcheck`, and `go vet`.
- Security scanning using `gosec`.
- Dockerized builds for multi-platform deployment (`linux/amd64`, `linux/arm64`).

## Requirements

- Go version 1.22 or later
- Docker
- `etcdctl` (optional, if interacting with etcd)

## Setup

### Cloning the Repository

Clone the repository:

```bash
git clone https://github.com/yourusername/etcdctl-go-decrypt.git
cd etcdctl-go-decrypt
```

### Running Locally

Ensure that you have Go installed (version 1.22 or later). Run the following command to build and run the project:

```bash
go build -o decrypt
./decrypt --key=<AES_KEY> --secret=<BASE64_ENCODED_SECRET>
```

### Running Tests

The project includes unit tests for AES decryption. Run the tests using:

```bash
go test ./...
```

## Development

### Makefile

The project includes a `Makefile` to automate tasks such as testing, building, and Docker operations. Below are some useful commands:

- **Run Static Analysis**

  Run static analysis with `golint`, `staticcheck`, and `go vet`:

  ```bash
  make static-analysis
  ```

- **Run Security Scan**

  Run `gosec` to check for security vulnerabilities:

  ```bash
  make security-scan
  ```

- **Build Docker Image**

  Build the Docker image and push it to DockerHub:

  ```bash
  make build
  ```

- **Clean Up**

  Remove the vendor directory:

  ```bash
  make clean
  ```

### Docker Build and Push

The project uses Docker for multi-platform builds (`linux/amd64`, `linux/arm64`). To build and push the Docker image to DockerHub, run:

```bash
make build
```

Ensure that you are logged in to DockerHub using:

```bash
docker login
```

The image will be tagged with the `latest`, version (based on the commit count), and Git commit SHA.

## Contributing

Feel free to open an issue or submit a pull request if you'd like to contribute.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

```

### Key Sections:
- **Features**: Lists the main features of the project.
- **Requirements**: Outlines dependencies such as Go and Docker.
- **Setup**: Provides instructions for running the project locally and running tests.
- **Development**: Describes how to use the `Makefile` to automate tasks.
- **CI/CD Pipeline**: Explains the GitHub Actions workflow for automated testing and Docker builds.
- **Contributing and License**: Encourages contributions and specifies the project’s license.

This `README.md` should give users a clear understanding of how to use and contribute to the project. Let me know if you need more customization!
