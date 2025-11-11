# Sample Go + React Starter

This repository provides a minimal full-stack scaffolding with a Go backend powered by Echo and a React + TypeScript frontend built on Vite and Tailwind CSS. It is intentionally lightweight so you can plug in additional routes, database models, and UI features while keeping the backend and frontend in the same tree.

## Highlights
- Echo-based HTTP API with logging and panic recovery middleware
- Database helper that can open either PostgreSQL or SQLite connections
- React UI wired with Tailwind CSS, shadcn/ui components, and path aliases
- Example health-check endpoint consumed by the frontend to demonstrate API integration
- Template files (`*.tmpl`) so you can rename the project and standardise ports

## Project Layout
```
.
├── main.go                 # Echo server wiring routes, middleware, and static assets
├── internal/
│   ├── db/db.go            # database/sql opener with PostgreSQL and SQLite drivers
│   ├── models/             # placeholder for domain models
│   └── routes/routes.go    # API route handlers
├── ui/                     # Vite + React frontend application
│   ├── src/App.tsx         # Example button invoking the health endpoint
│   ├── src/components/ui/  # shadcn/ui button implementation
│   ├── src/lib/utils.ts    # Tailwind utility helper
│   └── package.json.tmpl   # Template for the frontend package manifest
├── go copy.tmpl            # Helper template for generating an alternative go.mod
├── main.go.tmpl            # Template version of the server with tokenised ports
└── Makefile.tmpl           # Optional helper targets for build and run workflows
```

## Prerequisites
- Go 1.22 or newer
- Node.js 18+ (Vite works best with 18 or 20) and npm
- SQLite or PostgreSQL (only required if you plan to exercise the database helper)

## Configure the Templates
Several files include `{{ ... }}` placeholders so you can tailor the project name and ports:

1. **Backend port** – Update `main.go` to listen on a concrete port, e.g. `e.Start(":8080")`. Mirror the change in `main.go.tmpl` if you plan to reuse the template.
2. **Module name** – If you want a different module path for Go, copy `go copy.tmpl` over `go.mod` and replace `{{ .ProjectName }}` with your desired module (the committed `go.mod` already uses `module sample`).
3. **Frontend manifest** – Copy `ui/package.json.tmpl` to `ui/package.json` and replace `{{ .ProjectName }}` with the desired npm package name. The accompanying `ui/package-lock.json` still contains the placeholder string; feel free to regenerate it after running `npm install`.
4. **Vite config** – Copy `ui/vite.config.ts.tmpl` to `ui/vite.config.ts` and replace `{{ .FrontendPort }}` / `{{ .BackendPort }}` with the values you picked so the dev server can proxy API calls correctly.
5. **Document title** – `ui/index.html` contains `{{ .ProjectName }}` in the `<title>` tag; update it if you want a custom browser tab title.
6. **Makefile** – Rename `Makefile.tmpl` to `Makefile` to enable the predefined `make build-backend`, `make build-frontend`, and `make run-backend` targets.

Once the placeholders are replaced, the backend and frontend can run without additional templating infrastructure.

## Backend Setup
1. Export the database settings that `internal/db` expects:
   ```bash
   export DB_TYPE=sqlite3
   export DB_CONNECTION_STRING=./sample.db     # or a PostgreSQL DSN
   ```
2. Download dependencies (if needed) and start the server:
   ```bash
   go mod tidy
   go run main.go
   ```
   The server enables Echo's logger and recover middleware, stores the opened database connection on each request context (`c.Get("db")`), and serves files from `ui/dist` for the UI.

The API currently exposes a single endpoint:

| Method | Path             | Description                         |
| ------ | ---------------- | ----------------------------------- |
| GET    | `/api/v1/healthz` | Returns `{ "Result": true }` for health checks |

## Frontend Setup
After creating `ui/package.json` from the template:

```bash
cd ui
npm install
npm run dev    # starts Vite (defaults to http://localhost:5173)
```

The sample button in `src/App.tsx` issues a fetch to `/api/v1/healthz` and logs the response, demonstrating how the frontend talks to the backend. When you are ready to serve the UI from Go, build the production assets:

```bash
npm run build
```

The compiled files will be placed in `ui/dist` and automatically served by the Echo server.

## Database Notes
The helper in `internal/db` is intentionally thin—it wraps `database/sql` and relies on Go's standard driver registration. There is no migration tooling or ORM included, so bring your own (for example, `golang-migrate`, `sqlc`, or plain SQL files) and store related code under `internal/models` or another package of your choosing.

## Next Steps
- Flesh out additional routes inside `internal/routes` and mount them in `main.go`.
- Expand the frontend by adding components under `ui/src/components` and state management as needed.
- Add automated tests (`go test ./...` for the backend, Vitest/RTL for the frontend) and CI/check scripts as the project grows.
- Consider generating a real `Makefile` from `Makefile.tmpl` or wiring task runners to streamline local workflows.
