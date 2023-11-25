# Zocket Assignment


## Introduction
Designing and implementing a backend system using the Go programming
language for a product management application.


## Features
List out the key features of your application.

1- Create User / Fetch User

2- Create Product / Fetch Products

3- Image Analysis

4- Testing and Benchmarking


## Endpoints:

1. Fetching Details of all Users:
    
    - Endpoint: `http://localhost:3000/users/`
    - Method: GET

2. Creating a new User:

    - Endpoint: `http://localhost:3000/user/`
    - Method: POST
    - Request Body: User Details

    Example Request Body: 
    ```json
    {
        "name": "Naruto",
        "mobile": "9898989898",
        "latitude": 958.23,
        "longitude": 684.84
    }
    ```

3. Fetching Details of all Products:

    - Endpoint: `http://localhost:3000/products/`
    - Method: GET

4. Creating a new Product:

    - Endpoint: `http://localhost:3000/products/`
    - Method: POST
    - Request Body: Product Details

    Example Request Body:
    ```json
    {
        "user_id": 1,
        "product_name": "camera",
        "product_description": "Amazing Picture Taking Camera",
        "product_images": [
            "https://cdn.pixabay.com/photo/2017/01/09/14/16/camera-1966601_1280.jpg",
            "https://cdn.pixabay.com/photo/2017/01/09/14/16/camera-1966601_1280.jpg",
            "https://cdn.pixabay.com/photo/2017/01/09/14/16/camera-1966601_1280.jpg"
        ],
        "product_price": 26499
    }
    ```



## Installation & Getting Started

Detailed instructions on how to install, configure, and get the project running.

Clone the project

```bash
  https://github.com/Abhi0049k/Product-Management-Application.git
```

Install dependencies

```bash
  go mod tidy
```

create your own instance of database and migrate it

Run the Application

```bash
  go run main.go
```
