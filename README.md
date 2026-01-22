# AWS Secrets Manager Platform

A web platform to view and manage AWS Secrets Manager secrets with a React frontend and Go backend.

## Features

- **List Secrets**: View all secrets from AWS Secrets Manager with search/filter
- **View Secret Values**: Show/hide toggle for sensitive data
- **Create Secrets**: Add new secrets with name, value, and optional description
- **Update Secrets**: Modify existing secret values
- **Copy to Clipboard**: One-click copy for secret values
- **HTTP Basic Auth**: Secure access for local development
- **Swagger Documentation**: Interactive API documentation

## Project Structure

```
secrets-manager-platform/
├── backend/                          # Go REST API
│   ├── cmd/
│   │   └── server/
│   │       └── main.go               # Application entry point
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── secrets.go            # HTTP request handlers
│   │   ├── middleware/
│   │   │   └── auth.go               # Basic auth middleware
│   │   └── services/
│   │       └── secrets.go            # AWS Secrets Manager operations
│   ├── docs/                         # Swagger documentation (generated)
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── .env.example                  # Environment variables template
│   ├── go.mod
│   └── go.sum
├── frontend/                         # React application
│   ├── src/
│   │   ├── components/               # Reusable UI components
│   │   ├── routes/
│   │   │   ├── root.tsx              # Layout with navigation
│   │   │   ├── login.tsx             # Authentication page
│   │   │   ├── dashboard.tsx         # Secrets list view
│   │   │   ├── secret.tsx            # View/edit single secret
│   │   │   └── secret-new.tsx        # Create new secret form
│   │   ├── services/
│   │   │   └── api.ts                # API client with auth
│   │   ├── router.tsx                # React Router configuration
│   │   ├── main.tsx                  # Application entry point
│   │   └── index.css                 # Tailwind CSS imports
│   ├── package.json
│   └── vite.config.ts
├── plan/
│   └── aws-secrets-manager-platform-plan.md
├── Makefile                          # Commands to run the project
└── README.md
```

## Prerequisites

- **Go**: Version 1.21 or higher
- **Node.js**: Version 18 or higher
- **AWS Credentials**: Configured via one of:
  - Environment variables (`AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`)
  - AWS credentials file (`~/.aws/credentials`)
  - IAM role (when running on AWS infrastructure)

## Quick Start

### 1. Clone and Navigate

```bash
cd secrets-manager-platform
```

### 2. Configure Backend

```bash
cd backend

# Copy environment template
cp .env.example .env

# Edit .env with your settings
```

Edit the `.env` file:

```env
# Server Configuration
PORT=8080

# Authentication (change these for production!)
BASIC_AUTH_USER=admin
BASIC_AUTH_PASS=your-secure-password

# AWS Configuration
AWS_REGION=us-east-1
AWS_ACCESS_KEY_ID=your-access-key        # Optional if using IAM role
AWS_SECRET_ACCESS_KEY=your-secret-key    # Optional if using IAM role
```

### 3. Start Backend

```bash
# From the backend directory
go run cmd/server/main.go
```

The API server will start at `http://localhost:8080`

### 4. Configure and Start Frontend

Open a new terminal:

```bash
cd frontend

# Install dependencies
npm install

# Start development server
npm run dev
```

The frontend will start at `http://localhost:5173`

### 5. Access the Application

1. Open http://localhost:5173 in your browser
2. Log in with your configured credentials (default: `admin` / `admin`)
3. Start managing your AWS secrets

## Running Both Services

### Single Command (Recommended)

From the project root:
```bash
make dev
```

This starts both backend and frontend concurrently. Press `Ctrl+C` to stop both.

### Other Make Commands

```bash
make install   # Install all dependencies
make backend   # Run backend only
make frontend  # Run frontend only
make build     # Build both projects for production
make swagger   # Regenerate Swagger documentation
make setup     # Install deps + generate swagger
make help      # Show all available commands
```

### Manual (Separate Terminals)

**Terminal 1 - Backend:**
```bash
cd backend && go run cmd/server/main.go
```

**Terminal 2 - Frontend:**
```bash
cd frontend && npm run dev
```

## Swagger Documentation

Interactive API documentation is available via Swagger UI.

### Accessing Swagger UI

Once the backend is running, open your browser and navigate to:

```
http://localhost:8080/swagger/index.html
```

### Swagger Files

The API documentation is also available in these formats:
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **OpenAPI JSON**: http://localhost:8080/swagger/doc.json
- **Source files**: `backend/docs/swagger.yaml` and `backend/docs/swagger.json`

### Regenerating Swagger Documentation

If you modify the API, regenerate the documentation:

```bash
cd backend

# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init -g cmd/server/main.go -o docs
```

## API Reference

All endpoints require HTTP Basic Authentication.

### List Secrets

```http
GET /api/secrets
```

**Response:**
```json
[
  {
    "name": "my-secret",
    "description": "My application secret"
  }
]
```

### Get Secret Value

```http
GET /api/secrets/:name
```

**Response:**
```json
{
  "name": "my-secret",
  "value": "{\"key\": \"value\"}"
}
```

### Create Secret

```http
POST /api/secrets
Content-Type: application/json

{
  "name": "new-secret",
  "value": "{\"api_key\": \"abc123\"}",
  "description": "Optional description"
}
```

### Update Secret

```http
PUT /api/secrets/:name
Content-Type: application/json

{
  "value": "{\"api_key\": \"updated-value\"}"
}
```

## Tech Stack

### Frontend
- **React 19** - UI framework
- **Tailwind CSS v4** - Utility-first CSS
- **React Router v7** - Client-side routing with data loaders/actions
- **Axios** - HTTP client
- **Vite** - Build tool and dev server

### Backend
- **Go** - Programming language
- **Gin** - Web framework
- **AWS SDK for Go v2** - AWS Secrets Manager integration
- **Swaggo** - Swagger documentation generator
- **godotenv** - Environment variable management

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Backend server port | `8080` |
| `BASIC_AUTH_USER` | Authentication username | `admin` |
| `BASIC_AUTH_PASS` | Authentication password | `admin` |
| `AWS_REGION` | AWS region for Secrets Manager | - |
| `AWS_ACCESS_KEY_ID` | AWS access key (optional with IAM) | - |
| `AWS_SECRET_ACCESS_KEY` | AWS secret key (optional with IAM) | - |

## Production Build

### Backend
```bash
cd backend
go build -o server cmd/server/main.go
./server
```

### Frontend
```bash
cd frontend
npm run build
# Serve the dist/ directory with your preferred static file server
```

## Security Notes

- This application uses HTTP Basic Auth intended for local development
- Always use HTTPS in production environments
- Change default credentials before deployment
- Consider using a more robust authentication method for production use
