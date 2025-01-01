### **README for Train Ticket Booking System**

---

### **Project Overview**
This project is a REST API application for a train ticket booking system built using **Go**, **Gin Framework**, and **MongoDB**. The application provides basic authentication (register, login) and CRUD functionalities for managing trains and tickets.

---

### **Features**
- **User Authentication**: Register and login with JWT-based authentication.
- **Train Management**: Add and view train schedules.
- **Ticket Booking**: Book and cancel tickets.

---

### **Prerequisites**
Before setting up the project, ensure you have the following installed:
- GoLang
- `curl` for API testing (optional).

---

### **Project Structure**
```
trainTicketsGo/
├── main.go
├── routes/
│   ├── user.go
│   ├── train.go
│   └── ticket.go
├── controllers/
│   ├── authController.go
│   ├── trainController.go
│   └── ticketController.go
├── models/
│   ├── user.go
│   ├── train.go
│   └── ticket.go
├── database/
│   └── db.go
├── middlewares/
│   └── authMiddleware.go
├── go.sum
└── go.mod
```

---

### **Setup Instructions**

#### 1. Clone the Repository
```bash
git clone <repository_url>
cd trainTicketsGo
```

#### 2. Initialize Go Modules
```bash
go mod tidy
```

#### 3. Configure Environment Variables
Create a `.env` file in the root directory with the following content:
```env
DB_URI=mongodb://localhost:27017
DB_NAME=trainTicketGo
JWT_SECRET=your_jwt_secret
```

#### 4. Run the Application
Start the server with the following command:
```bash
go run main.go
```
The server will start at `http://localhost:3000`.

---

### **API Endpoints**

#### **Public Routes**
1. **Register**:
   - **URL**: `api/v1/auth/register`
   - **Method**: `POST`
   - **Payload**:
     ```json
     {
       "name": "sumit",
       "email": "sumit@bhuia.com",
       "password": "1234567"
     }
     ```
   - **cURL**:
     ```bash
     curl -X POST http://localhost:3000/api/v1/auth/register \
     -H "Content-Type: application/json" \
     -d '{
         "name": "sumit",
         "email": "sumit@bhuia.com",
         "password": "1234567"
     }'
     ```

2. **Login**:
   - **URL**: `api/v1/auth/login`
   - **Method**: `POST`
   - **Payload**:
     ```json
     {
       "email": "sumit@bhuia.com",
       "password": "1234567"
     }
     ```
   - **cURL**:
     ```bash
     curl -X POST http://localhost:3000/api/v1/auth/login \
     -H "Content-Type: application/json" \
     -d '{
         "email": "sumit@bhuia.com",
         "password": "1234567"
     }'
     ```

#### **Protected Routes (JWT Required)**

1. **Get Trains**:
   - **URL**: `api/v1/trains`
   - **Method**: `GET`
   - **cURL**:
     ```bash
     curl -X GET http://localhost:3000/api/v1/trains \
     -H "Authorization: Bearer <your_jwt_token>"
     ```

2. **Add Train**:
   - **URL**: `api/v1/trains`
   - **Method**: `POST`
   - **Payload**:
     ```json
     {
       "name": "Express Train",
       "source": "City A",
       "destination": "City B",
       "schedule": "2024-01-01 10:00:00"
     }
     ```
   - **cURL**:
     ```bash
     curl -X POST http://localhost:3000/api/v1/trains \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <your_jwt_token>" \
     -d '{
         "name": "Express Train",
         "source": "City A",
         "destination": "City B",
         "schedule": "2024-01-01 10:00:00"
     }'
     ```

3. **Book Ticket**:
   - **URL**: `api/v1/tickets`
   - **Method**: `POST`
   - **Payload**:
     ```json
     {
       "user_id": "64ff3c2b7b18a3d47e38e1b9",
       "train_id": "64ff3c5d7b18a3d47e38e1bc",
       "seat": "A1-23"
     }
     ```
   - **cURL**:
     ```bash
     curl -X POST http://localhost:3000/api/v1/tickets \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <your_jwt_token>" \
     -d '{
         "user_id": "64ff3c2b7b18a3d47e38e1b9",
         "train_id": "64ff3c5d7b18a3d47e38e1bc",
         "seat": "A1-23"
     }'
     ```

4. **Cancel Ticket**:
   - **URL**: `api/v1/tickets/cancel`
   - **Method**: `POST`
   - **Payload**:
     ```json
     {
       "ticket_id": "64ff3c7e7b18a3d47e38e1bf"
     }
     ```
   - **cURL**:
     ```bash
     curl -X POST http://localhost:3000/api/v1/tickets/cancel \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <your_jwt_token>" \
     -d '{
         "ticket_id": "64ff3c7e7b18a3d47e38e1bf"
     }'
     ```

---

### **Testing Instructions**
1. Use the `curl` commands provided above or tools like [Postman](https://www.postman.com/) to test API endpoints.
2. Ensure you have a valid JWT token for protected routes. You can get it by logging in using the `api/v1/auth/login` endpoint.

