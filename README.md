# Appointy_API

Tech Stack
Language: Golang
Web Framework: Gin
Database: MongoDB
Why this tech stack was chosen
Golang is a highly performant and efficient language that excels in building scalable and concurrent systems. It also has a rich standard library that makes it easy to implement HTTP servers and handle network requests. Gin is a lightweight and fast web framework that is built on top of the net/http package and provides helpful features like routing, middleware, and error handling. MongoDB was chosen as the database because of its flexible schema and ability to handle large amounts of data, making it a good fit for a system that needs to store and retrieve locations of 50,000 drivers.

Infrastructure Requirements

Golang: Version 1.16 or higher
MongoDB: Version 4.0 or higher
Operating System: Linux, macOS or Windows
Setup Instructions
Clone the repository: git clone https://github.com/<username>/driver-location
Install Golang and MongoDB if not already installed
Navigate to the project directory and run the command go get to install the necessary dependencies
Run MongoDB using the command mongod
To run the application, run the command go run main.go
The server will start running on localhost:8080
Automated Deployment to Development and Test Environment:

To automate the deployment process, you can use containerization tools like Docker and orchestration tools like Kubernetes or Docker Compose.
Dockerize the application by creating a Dockerfile that specifies the application and its dependencies.
Use Kubernetes or Docker Compose to deploy the application to development and test environments.
Automate the deployment process using a Continuous Integration and Continuous Deployment (CI/CD) tool like Jenkins or GitLab CI/CD.