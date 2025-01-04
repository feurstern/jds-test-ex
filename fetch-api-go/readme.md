# Fetch Api Project : GO-BLOG

This README provides detailed instructions to install, build, and run the backend project using Docker and locally.

---

## Prerequisites

Before starting, ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/) (version 20.10 or higher recommended)
- [Docker Compose](https://docs.docker.com/compose/install/) (version 1.29 or higher recommended)
- Go (if running locally) - [Download Go](https://golang.org/dl/)

---

## Environment Variables

The project uses an `.env` file for environment variables. Below is an example of what your `.env` file should look like:

```
JWT_KEY_SECRET=Yukirinkawaii123!
LOCALHOST =localhost:8181
CLIENT_HOST = http://localhost:9000
API_URL= https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product
USD_TO_IDR = 16031
```

Ensure the `.env` file is located in the project root directory.

---

## Run with Docker

### Step 1: Build the Docker Image

1. Navigate to the project directory:
   ```bash
   cd fetch-api-go
   ```

2. Build the Docker image:
   ```bash
   docker-compose build
   ```

### Step 2: Start the Docker Container

1. Run the container:
   ```bash
   docker-compose up
   ```

2. The application will be available at:
   ```
   http://localhost:8181
   ```

### Step 3: Stop the Docker Container

1. To stop the container, press `CTRL+C` in the terminal running Docker Compose.
2. Alternatively, you can stop the containers with:
   ```bash
   docker-compose down
   ```

---

## Run Locally

### Step 1: Install Dependencies

1. Navigate to the project directory:
   ```bash
   cd fetch/api
   ```

2. Install Go modules:
   ```bash
   go mod download
   ```

### Step 2: Build the Application

1. Build the Go binary:
   ```bash
   go build -o main .
   ```

### Step 3: Run the Application

1. Start the application:
   ```bash
   ./main
   ```

2. The application will be available at:
   ```
   http://localhost:8181
   ```

---

## Troubleshooting

1. **Docker Issues**
   - Ensure your `.env` file is correctly placed and contains all required variables.
   - Check if port `8181` is already in use. If so, stop any services using it or change the port in the `docker-compose.yml` file.

2. **Local Build Issues**
   - Ensure you have Go installed and the required version is being used.
   - Use the following command if cross-compilation issues occur:
     ```bash
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
     ```

---

## Project Structure

```
/fetch-api-go
 ├── main.go         # Entry point of the application
 ├── Dockerfile     # Dockerfile for building the container
 ├── docker-compose.yml  # Docker Compose configuration
 ├── go.mod         # Go module dependencies
 ├── .env           # Environment variables
```

---

## License

This project is licensed under the [MIT License](LICENSE).

