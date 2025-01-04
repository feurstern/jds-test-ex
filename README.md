# authentication-app-express

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
PORT=5454
JWT_KEY_SECRET=Yukirinkawaii123
```


To run with docker:
I am asssume that you're currently using Linux, and my current machine is using arch linux

```bash
cd authentication-app-express
```

```bash
sudo docker compose up
```

Ensure the `.env` file is located in the project root directory.

---



## Run with Docker

### Step 1: Build the Docker Image

1. Navigate to the project directory:
   ```bash
   cd authentication-app-express
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
   http://localhost:5454
   ```

### Step 3: Stop the Docker Container

1. To stop the container, press `CTRL+C` in the terminal running Docker Compose.
2. Alternatively, you can stop the containers with:
   ```bash
   docker-compose down
   ```

---


# Run Locally

### Step 1: Install Dependencies

1. Navigate to the project directory:
   ```bash
   cd  authentication-app-express
   ```

2. Install bun modules:
   ```bash
   bun i
   ```

### Step 2: Run the Application

1. Start the application:
   ```bash
   bun.index
   ```

2. The application will be available at:
   ```
   http://localhost:5454
   ```

---
# Troubleshooting

1. **Docker Issues**
   - Ensure your `.env` file is correctly placed and contains all required variables.
   - Check if port `8181` is already in use. If so, stop any services using it or change the port in the `docker-compose.yml` file.

2. **Local Build Issues**
   - Ensure you have Go installed and the required version is being used.

---



![image](https://github.com/user-attachments/assets/3f4df0c3-b86c-4734-8976-4c9076a6052b)

// end if installation

---
## API :

1. **Create user**
   Post ->  http://localhost:5454/auth/create

![image](https://github.com/user-attachments/assets/c3f359c8-822b-4b17-8da9-1ebab54af46e)


Copy paste the password from the terminal in order to test the next API

![image](https://github.com/user-attachments/assets/d35ddd95-1081-4eb7-99aa-c40d1d55ba51)


2. **Login**
   Post -> http://localhost:5454/auth/login
![image](https://github.com/user-attachments/assets/79951119-6994-4a01-ad84-d226bdf40a19)

3. **Private-claim**
   Post ->  http://localhost:5454/auth/private

   Ensure that you already copy paste the given token from the previous phase.
![image](https://github.com/user-attachments/assets/2d29495d-8e20-4901-88d8-0af84e978e12)

---

## Project Structure

```
/authentication-app-express
 ├── auth
     ├── claim.js
     ├── createUser.js
     ├── index.js
     ├── login.js
 ├── helper
     ├── utils.js
 ├── test
      ├── auth.test.js
 ├── Dockerfile     
 ├── docker-compose.yml  # Docker Compose configuration
 ├── package.json        
 ├── .env           # Environment variables
 ├── jsconfi.json
```

---





