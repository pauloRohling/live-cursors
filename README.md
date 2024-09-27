# Live Cursors

Live Cursors is a simple web application that allows you to share your cursors with others.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Node.js](https://nodejs.org/en/download/)
- [Air](https://github.com/air-verse/air)
- [API Ninjas Account](https://api-ninjas.com/)

### Installation

1. Create an account on [API Ninjas](https://api-ninjas.com/) and get your API key.

2. Clone the repository:

```bash
git clone https://github.com/pauloRohling/live-cursors.git
```

3. Navigate to the project directory:

```bash
cd live-cursors
```

4. Build the Docker image:

```bash
docker build -t live-cursors .
```

5. Run the Docker container:

```bash
docker run --rm -e API_KEY=<YOUR_API_KEY> -p 8080:8080 live-cursors
```

6. Navigate to the web project directory:

```bash
cd web
```

7. Install the dependencies:

```bash
npm install
```

8. Start the development server:

```bash
npm start
```

9. Open your web browser and navigate to `http://localhost:4200`.

## Environment Variables

| Variable                 | Description                                 | Default Value                            | Required |
|--------------------------|---------------------------------------------|------------------------------------------|----------|
| `SERVER_PORT`            | Port to listen on                           | 8080                                     | false    |
| `API_URL`                | URL of the API                              | https://api.api-ninjas.com/v1/randomuser | false    |
| `API_KEY`                | API key for the API                         | -                                        | true     |
| `HTTP_MAX_RETRY`         | Maximum number of retries for HTTP requests | 3                                        | false    |
| `HTTP_MAX_RETRY_TIMEOUT` | Maximum timeout for HTTP requests           | 5s                                       | false    |
| `HTTP_MIN_RETRY_TIMEOUT` | Minimum timeout for HTTP requests           | 1s                                       | false    |

## License

Live Cursors is released under the MIT License. See the [LICENSE](LICENSE) file for more information.