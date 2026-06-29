# Go Shop API - E-Commerce Backend

A complete e-commerce backend API built with Go, Gin, GORM, and PostgreSQL.

## ✨ Features

### 🔐 Authentication & Authorization
- User registration with email validation
- User login with JWT token authentication
- Token refresh mechanism for session management
- User logout functionality
- Role-based access control (Admin/User roles)
- JWT middleware for protected routes

### 👤 User Management
- User profile retrieval
- User profile updates
- Secure password hashing
- User data persistence

### 🏪 Category Management
- Create product categories (Admin only)
- Update category information (Admin only)
- Delete categories (Admin only)
- View all categories (Public)

### 📦 Product Management
- Create products (Admin only)
- Update product details (Admin only)
- Delete products (Admin only)
- Get all products (Public)
- Get single product details (Public)
- Product search and filtering

### 🖼️ Product Images
- Upload product images (Admin only)
- Manage multiple images per product
- Image metadata (alt text, timestamps)
- Local and S3 storage support
- Image deletion

### 🛒 Shopping Cart
- Add items to cart
- Update cart item quantities
- Remove items from cart
- View cart contents
- Cart persistence per user

### 📋 Order Management
- Create orders from cart
- Order history tracking
- Order status management
- Order details and items
- Order payment processing

### 📤 File Upload
- Local file storage provider
- AWS S3 integration support
- Configurable upload location
- Image validation and management

### 🔧 Infrastructure
- PostgreSQL database with migrations
- GORM ORM for database operations
- Gin web framework
- CORS middleware
- Request logging
- Error handling and recovery
- Docker & Docker Compose setup
- Nginx reverse proxy configuration
- LocalStack for AWS S3 emulation

## 🏗️ Database Schema

### Tables
- **users** - User accounts and authentication
- **refresh_tokens** - JWT refresh token management
- **categories** - Product categories
- **products** - Product information
- **product_images** - Product images with alt text and timestamps
- **carts** - Shopping carts per user
- **cart_items** - Items in shopping cart
- **orders** - Customer orders
- **order_items** - Items in each order

## 📡 API Endpoints

### Authentication (Public)
```
POST   /api/v1/auth/register      - User registration
POST   /api/v1/auth/login         - User login
POST   /api/v1/auth/refresh       - Refresh JWT token
POST   /api/v1/auth/logout        - User logout
```

### User Management (Protected)
```
GET    /api/v1/users/profile      - Get user profile
PUT    /api/v1/users/profile      - Update user profile
```

### Categories
```
GET    /api/v1/categories         - Get all categories (Public)
POST   /api/v1/categories         - Create category (Admin)
PUT    /api/v1/categories/:id     - Update category (Admin)
DELETE /api/v1/categories/:id     - Delete category (Admin)
```

### Products
```
GET    /api/v1/products           - Get all products (Public)
GET    /api/v1/products/:id       - Get product details (Public)
POST   /api/v1/products           - Create product (Admin)
PUT    /api/v1/products/:id       - Update product (Admin)
DELETE /api/v1/products/:id       - Delete product (Admin)
POST   /api/v1/products/:id/images - Upload product image (Admin)
```

### Shopping Cart (Protected)
```
GET    /api/v1/carts              - Get cart
POST   /api/v1/carts/items        - Add to cart
PUT    /api/v1/carts/items/:id    - Update cart item
DELETE /api/v1/carts/items/:id    - Remove from cart
```

### Orders (Protected)
```
POST   /api/v1/orders             - Create order
GET    /api/v1/orders             - Get order history
GET    /api/v1/orders/:id         - Get order details
```

## 🛠️ Tech Stack

- **Language**: Go 1.x
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT
- **API Documentation**: REST
- **Containerization**: Docker & Docker Compose
- **Reverse Proxy**: Nginx
- **Storage**: Local File System & AWS S3 (via LocalStack)

## 🗂️ Project Structure

```
├── cmd/
│   └── api/
│       └── main.go              - Application entry point
├── db/
│   └── migrations/              - Database migrations
├── docker/
│   ├── docker-compose.yml       - Container orchestration
│   ├── nginx.conf               - Nginx configuration
│   └── init/                    - Initialization scripts
├── internal/
│   ├── config/                  - Configuration management
│   ├── database/                - Database connection
│   ├── dto/                     - Data transfer objects
│   ├── interfaces/              - Interface definitions
│   ├── logger/                  - Logging setup
│   ├── models/                  - Data models
│   ├── provider/                - Upload providers
│   ├── server/                  - HTTP server & handlers
│   ├── services/                - Business logic
│   └── utils/                   - Utility functions
└── uploads/                     - Local file storage
```

## 🚀 Getting Started

### Prerequisites
- Go 1.x
- Docker & Docker Compose
- PostgreSQL (or use Docker)

### Installation

1. Clone the repository
```bash
git clone <repository-url>
cd go-shop-api
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run database migrations
```bash
make migrate-up
```

5. Start the application
```bash
go run cmd/api/main.go
```

### Using Docker

```bash
docker-compose -f docker/docker-compose.yml up -d
```

## 📝 Development

### Run Tests
```bash
make test
```

### Build Binary
```bash
make build
```

### Linting
```bash
make lint
```

## 🔒 Security Features

- JWT-based authentication
- Password hashing with bcrypt
- CORS middleware
- Admin role verification
- Protected API routes
- Secure token refresh mechanism

## 📚 API Documentation

For detailed API documentation and examples, please refer to the endpoint descriptions above.

## 🤝 Contributing

1. Create a feature branch
2. Commit your changes
3. Push to the branch
4. Create a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👨‍💻 Author

GALANG Development Team
