# Sample Go + React Starter

This repository packages a minimal full-stack starter that combines a Go 1.22 API (built with Echo) and a Vite-powered React + TypeScript frontend. It is intended as a launch point for experimenting with an `echo` backend, shadcn/ui components, and a small amount of database wiring that can target either PostgreSQL or SQLite.

## Highlights
- Go service with logging/recovery middleware and a simple health endpoint
- Small database abstraction that can open either PostgreSQL or SQLite connections via environment variables
- React 18 frontend bootstrapped with Vite, Tailwind CSS, tailwind-merge, and shadcn/ui components
- Production build pipeline that serves the compiled frontend assets from the Go process
- Scaffolding templates (`*.tmpl`) for stamping project-specific names, ports, and build steps

## Repository Layout
```
.
├── main.go                  // Go entrypoint (templated version in main.go.tmpl)
├── internal/
│   ├── db/db.go             // Database factory returning a shared *sql.DB
│   └── routes/routes.go     // Echo handlers (currently only /api/v1/healthz)
├── ui/                      // React + Vite frontend
│   ├── src/App.tsx          // Demo button calling the health check
│   ├── src/components/      // shadcn/ui powered button
│   └── tailwind.config.js   // Tailwind + shadcn configuration
├── go.mod / go.sum          // Backend dependencies (templated copy in go copy.tmpl)
└── Makefile.tmpl            // Optional build helpers for backend/frontend
```

## Prerequisites
- Go 1.22.3+
- Node.js 18+ with npm
- Either SQLite (bundled) or a reachable PostgreSQL instance

## Configure Project Metadata
Several files (`main.go.tmpl`, `go copy.tmpl`, `Makefile.tmpl`, `ui/vite.config.ts.tmpl`, `ui/package.json.tmpl`) contain template markers such as `{{ .ProjectName }}` and `{{ .BackendPort }}`. Copy or process these templates to inject the values that match your project. A quick manual setup looks like:
```bash
cp go\ copy.tmpl go.mod        # Replace {{ .ProjectName }} with your module path
cp main.go.tmpl main.go        # Replace {{ .BackendPort }} with a port (e.g. 8080)
cp ui/package.json.tmpl ui/package.json
cp ui/vite.config.ts.tmpl ui/vite.config.ts
```
If you are wiring this repository into a generator, hook those files into your templating step instead.

## Backend Setup
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Export the database configuration:
   ```bash
   export DB_TYPE=sqlite3
   export DB_CONNECTION_STRING=./db.sqlite
   # For PostgreSQL:
   # export DB_TYPE=postgres
   # export DB_CONNECTION_STRING=postgres://user:pass@localhost:5432/app?sslmode=disable
   ```
3. Run the API (replace `8080` with whatever port you templated in `main.go`):
   ```bash
   go run main.go
   ```
   Echo will mount the API under `/api/v1/*` and serve any built frontend assets from `ui/dist`.

## Frontend Setup
1. Create the non-templated configuration files if you have not already:
   ```bash
   cp ui/package.json.tmpl ui/package.json
   cp ui/vite.config.ts.tmpl ui/vite.config.ts
   ```
2. Install and build:
   ```bash
   cd ui
   npm install
   npm run build   # emits ui/dist for the Go server to serve
   ```
3. Run the Vite dev server during development (proxied requests to `/api` will hit the backend if you configure Vite accordingly):
   ```bash
   npm run dev
   ```

## API Surface
- `GET /api/v1/healthz` — Returns `{"Result": true}` when the server and database plumbing are reachable. The default frontend button in `App.tsx` triggers this endpoint and logs the result to the browser console.

## Development Tips
- Echo automatically recovers from panics and logs structured requests; use `ECHO_LOG_LEVEL=DEBUG` for verbose logs.
- When iterating on the frontend, run `npm run dev` from `ui/` and start the Go server separately; configure Vite's proxy in `ui/vite.config.ts` to forward `/api` to your API port.
- The `Makefile.tmpl` provides convenience targets (`build-backend`, `build-frontend`, `run-backend`) once you stamp in your project name. Copy + adjust it if you prefer using make.
- Add new API routes under `internal/routes` and register them in `main.go`; the `db` connection is injected into `echo.Context` for each request via middleware.

## License
This starter is provided as-is for experimentation and internal projects. Adapt, extend, and redistribute as your use case requires.
