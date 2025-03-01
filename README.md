# Live Cursors

Live Cursors is a sample web application that allows you to share your cursors with others.

This project was created as an experiment about how to use WebSockets to implement a real-time cursor sharing application. The application uses [Gorilla WebSocket](https://github.com/gorilla/websocket) to send and receive cursor positions in real-time.

## Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [API Ninjas Account](https://api-ninjas.com/)
- [Node.js](https://nodejs.org/en/download/)
- [Air](https://github.com/air-verse/air)
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

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