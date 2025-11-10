# Sample Go Web Application

A full-stack web application built with Go (Echo framework) on the backend and React with TypeScript on the frontend.

## Architecture

This project consists of two main components:

- **Backend**: Go server using the Echo web framework
- **Frontend**: React application with TypeScript and Vite

## Features

- RESTful API with health check endpoint
- Database abstraction layer supporting multiple databases (PostgreSQL, SQLite)
- React frontend with Tailwind CSS and shadcn/ui components
- Middleware for logging and error recovery

## Tech Stack

### Backend
- [Go](https://golang.org/) 1.22.3
- [Echo](https://echo.labstack.com/) v4 - Web framework
- [PostgreSQL](https://www.postgresql.org/) / [SQLite](https://www.sqlite.org/) - Database support
- Environment-based configuration

### Frontend
- [React](https://react.dev/) with TypeScript
- [Vite](https://vitejs.dev/) - Build tool
- [Tailwind CSS](https://tailwindcss.com/) - Styling
- [shadcn/ui](https://ui.shadcn.com/) - UI components

## Project Structure

```
.
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── internal/
│   ├── db/                 # Database abstraction layer
│   │   └── db.go
│   └── routes/             # API route handlers
│       └── routes.go
└── ui/                     # Frontend React application
    ├── src/
    │   ├── App.tsx         # Main React component
    │   ├── main.tsx        # React entry point
    │   └── components/     # React components
    └── package.json
```

## Prerequisites

- Go 1.22.3 or higher
- Node.js and npm (for frontend development)
- PostgreSQL or SQLite (depending on your configuration)

## Setup

### Environment Variables

Create a `.env` file or set the following environment variables:

```bash
DB_TYPE=sqlite3                    # or "postgres"
DB_CONNECTION_STRING=./db.sqlite   # or PostgreSQL connection string
```

For PostgreSQL:
```bash
DB_TYPE=postgres
DB_CONNECTION_STRING=postgres://user:password@localhost:5432/dbname?sslmode=disable
```

### Backend Setup

1. Install Go dependencies:
```bash
go mod download
```

2. Build the application:
```bash
go build -o app
```

3. Run the application:
```bash
export DB_TYPE=sqlite3
export DB_CONNECTION_STRING=./db.sqlite
./app
```

Or run directly:
```bash
go run main.go
```

### Frontend Setup

1. Navigate to the UI directory:
```bash
cd ui
```

2. Install dependencies:
```bash
npm install
```

3. For development (with hot reload):
```bash
npm run dev
```

4. Build for production:
```bash
npm run build
```

The built files will be placed in `ui/dist/` and served by the Go backend.

## API Endpoints

### Health Check
```
GET /api/v1/healthz
```

Returns the health status of the application.

**Response:**
```json
{
  "Result": true
}
```

## Development

### Running the Full Stack

1. Build the frontend:
```bash
cd ui && npm run build && cd ..
```

2. Start the backend (which will serve the frontend):
```bash
go run main.go
```

The application will be available at `http://localhost:8080` (default port may vary based on configuration).

### Database

The application uses a database abstraction layer that supports:
- SQLite (for development/testing)
- PostgreSQL (for production)

Configure via environment variables as shown in the Setup section.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is provided as-is for educational and development purposes.
