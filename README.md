# AWS Serverless CRUD API in Golang

This repository contains a full serverless stack CRUD API built with Golang, leveraging AWS services such as API Gateway, Lambda, and DynamoDB. The API allows for Create, Read, Update, and Delete operations on a simple User table.

## Features
- **API Gateway:** Handles HTTP requests and routes them to the appropriate Lambda functions.
- **Lambda**: Executes the backend logic in a serverless environment.
- **DynamoDB**: Provides a scalable, NoSQL database for storing resource data.
- **Golang**: The API is written in Golang, ensuring fast and efficient execution.

## Getting Started
### Prerequisites
- AWS account with access to API Gateway, Lambda, and DynamoDB.
- Golang installed on your local machine.
- AWS CLI configured with appropriate credentials.

### Deployment
1. **Clone the Repository:**
   ```powershell
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name
   ```

2. **Setup AWS Services:**
  - Create a DynamoDB table with a primary key as `email`.
  - Set up API Gateway with routes for each CRUD operation.
  - Create corresponding Lambda functions and link them to API Gateway routes.

3. **Create a zip from binary:**
    - Before deploying the Lambda function, you need to create a ZIP file from your compiled Go binary.
    - **Steps to create the ZIP file,**
   
        1. **Build the Binary**:
         - Compile your Go code into a binary that is compatible with AWS Lambda. For example,
            ```powershell
              GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
            ```
         - don't forget to name the binary executable as `bootstrap`, it is mandatory.
           
        2. **Package the Binary**:
          - Create a ZIP file containing the compiled binary
           ```powershell
             zip bootstrap.zip main
           ```
        - This ZIP file will be uploaded to AWS Lambda via the AWS Management Console.
  

5. **Deploy Lambda Functions Using AWS UI:**
   - Go to the AWS Management Console.
   - Navigate to Lambda and click on the Create function.
   - Choose an Author from scratch and provide a function name.
   - Select Golang as the runtime.
   - Upload your main.go file as a zip package under the Function code section.
   - Configure the function's execution role and other settings as required.
   - Repeat these steps for each CRUD operation (Create, Read, Update, Delete).
  
### **API Endpoints: After deployment, you can access the following API endpoints:**
  - Create user: `POST` `/<apigateway_url> --data <resource>`
  - Get user: `GET` `/<apigateway_url>/\?email\=`
  - Get all users: `GET` `/<apigateway_url>`
  - Update user: `PUT` `/<apigateway_url> --data <resource>`
  - Delete user: `DELETE` `/<apigateway_url>/\?email\=`
  
### **Example Usage**
  Here's how to create a new resource using the API:
   - Create user:
     ```powershell
        curl --location --request POST 'https://ly7abx6bgi.execute-api.ap-south-1.amazonaws.com/staging' \
         --data '{ "email": "rohan@gmail.com", "firstname": "rohan", "lastname": "yh" }'
     ```

   - Get user:
     ```powershell
       curl -X GET https://ly7abx6bgi.execute-api.ap-south-1.amazonaws.com/staging\?email\=rohan@gmail.com
     ```

   - Get all users:
     ```powershell
       curl -X GET https://ly7abx6bgi.execute-api.ap-south-1.amazonaws.com/staging
     ```

  - Update user:
      ```powershell
         curl --location --request PUT https://ly7abx6bgi.execute-api.ap-south-1.amazonaws.com/staging \
         --data '{ "email": "rohan@gmail.com", "firstname": "rohan", "lastname": "honnakatti" }'
       ```

- Delete user:
  ```powershell
    curl -X DELETE https://ly7abx6bgi.execute-api.ap-south-1.amazonaws.com/staging\?email\=rohan@gmail.com
  ```

### Contributing
Contributions are welcome! If you have suggestions for improving this project or would like to contribute code, please feel free to fork the repository and submit a pull request. Make sure to follow the existing code style and include relevant tests with your submission.


### Contact,
If you have any questions, or feedback, or need further assistance, feel free to reach out. You can contact me at rohanyh101@gmail.com.
