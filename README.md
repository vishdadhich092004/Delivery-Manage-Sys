# Delivery Management System

Delivery Management System is a web application designed to streamline the process of managing delivery orders. The system allows users to input order details, track delivery statuses, and manage delivery personnel, ensuring smooth and efficient operations.

## Features

- **Order Management:** Add, update, and track delivery orders.
- **Delivery Personnel Tracking:** Monitor the status and progress of delivery personnel.
- **Real-time Updates:** Instant updates on the status of orders and deliveries.
- **User Management:** Manage users who interact with the system.

## Technologies Used

- **Frontend:** Vue
- **Backend:** Golang with Gin framework
- **Database:** PostgreSQL

## Installation

### Backend (Golang)

1. Clone the repository:

    ```bash
    git clone https://github.com/vishdadhich092004/Delivery-Manage-Sys.git
    cd Delivery-Manage-Sys
    ```
2. Navigate to the server directory:

    ```bash
    cd server
    ```

3. Install Go dependencies:

    ```bash
    go mod tidy
    ```

4. Run the backend:

    ```bash
    go run .\cmd\main.go
    ```

### Frontend (Vue)   

1. Navigate to the client directory:

    ```bash
    cd client
    ```

2. Install frontend dependencies:

    ```bash
    npm install
    ```

3. Start the frontend server:

    ```bash
    npm run dev
    ```


