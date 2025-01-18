# Weather API

This is a weather API project that fetches and returns weather data from a 3rd party API. This project is based on the [Weather API Project Idea](https://roadmap.sh/projects/weather-api-wrapper-service) from the [roadmap.sh](https://roadmap.sh) platform.

## Project Overview

Instead of relying on our own weather data, this project aims to build a weather API that fetches and returns weather data from a 3rd party API. This project will help you understand how to work with 3rd party APIs, caching, and environment variables.

The project uses the [Visual Crossing's Weather API](https://www.visualcrossing.com/weather-api) as the 3rd party API to fetch weather data. This API is completely free and easy to use.

For caching the weather data, the project uses Redis, a popular in-memory data structure store. The city code entered by the user is used as the key, and the result from calling the 3rd party API is cached with a configurable expiration time (e.g., 12 hours).

The project also includes Swagger documentation for the API, making it easier to understand and interact with the endpoints.

## Technologies Used

- Programming Language: [Go](https://golang.org/)
- Web Framework: [Gin](https://github.com/gin-gonic/gin)
- Caching: [Redis](https://redis.io/)
- API Documentation: [Swagger](https://swagger.io/)
- Docker: [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)

## Features

- Fetch weather data from a 3rd party API (Visual Crossing's Weather API)
- Cache the weather data using Redis to improve performance
- Implement environment variables to store sensitive information like API keys and Redis connection details
- Handle errors properly, such as when the 3rd party API is down or the city code is invalid
- Implement rate limiting to prevent abuse of the API
- Provide Swagger documentation for the API

## Getting Started

1. Clone the repository:

```
git clone https://github.com/alielmi98/go-weather-api.git
```

2. Build and run the application using Docker Compose:

```
docker-compose build
docker-compose up
```

This will build the Docker images and start the application and Redis containers.

3. The Weather API will be available at `http://localhost:8080`.

## Contribution

This project is part of the [roadmap.sh](https://roadmap.sh) platform, which is designed to help developers learn and grow their skills. If you find any issues or have suggestions for improvement, please feel free to contribute to the project by submitting a pull request or opening an issue.

## License

This project is licensed under the [MIT License](LICENSE).