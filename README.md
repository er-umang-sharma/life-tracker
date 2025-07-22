# Life Tracker

Life Tracker is a modular application for tracking various aspects of daily life. It consists of a Golang microservice backend and a modern front-end (React/Next.js recommended).

## Features
- Daily Habits Tracker
- Food Tracker
- Weights Tracker
- Daily Goals Tracker
- Daily Journal

## Project Structure

```
life-tracker/
├── backend/         # Golang microservices
│   ├── cmd/        # Main entrypoints for each service (habits, food, etc.)
│   ├── internal/   # Service logic and business code
│   ├── api/        # API contracts (OpenAPI/Protobuf)
│   └── go.mod      # Go module definition
├── frontend/        # Front-end app (React/Next.js)
│   ├── public/
│   ├── src/
│   └── ...
├── README.md
└── ...
```

## Backend Setup (Habits Microservice Example)
1. Ensure you have Go and PostgreSQL installed.
2. Create a `.env` file in `backend/` with:
   ```
   DATABASE_URL=postgres://username:password@localhost:5432/life-tracker
   ```
3. Create the `habits` table:
   ```sql
   CREATE TABLE habits (
     id SERIAL PRIMARY KEY,
     name TEXT NOT NULL
   );
   ```
4. Install dependencies:
   ```sh
   go mod tidy
   ```
5. Run the habits service:
   ```sh
   go run ./cmd/habits
   ```

## Frontend Setup
1. Navigate to `frontend/` and install dependencies:
   ```sh
   npm install
   ```
2. Start the development server:
   ```sh
   npm start
   ```

## Contributing
Feel free to open issues or submit PRs for new features or improvements.