# Ads Impression Counter Service

This repository contains an **in-memory concurrent service** for tracking ad impressions in real-time. The service provides:

- **Campaign Management**: Register and manage advertising campaigns.
- **Impression Tracking**: Log unique ad impressions while avoiding duplicates using a TTL mechanism.
- **Statistics Reporting**: Retrieve impression counts for the last hour, last day, and total impressions.
- **REST API Exposure**: Standard HTTP endpoints for campaign management and analytics.
- **Thread-Safe Concurrency**: Ensures safe concurrent operations using Go's `sync.Mutex`.

---

## Features

### **1. Campaign Management**

- Create campaigns with a `name` and `start_time`.
- Persist campaigns in-memory.

### **2. Impression Tracking**

- Log user impressions for a given `campaign_id`, `user_id`, and `ad_id`.
- Prevent duplicate impressions within a **1-hour TTL** window.

### **3. Campaign Statistics**

- Retrieve aggregated impression stats.
- Supported statistics:
   - **Last Hour**
   - **Last Day**
   - **Total Count**

### **4. API Endpoints**

- `POST /api/v1/campaigns` — Create a campaign
- `POST /api/v1/impressions` — Track an impression
- `GET /api/v1/campaigns/{id}/stats` — Get campaign stats
- `404` handling for invalid routes

### **5. Concurrent & Thread-Safe**

- Uses `sync.Mutex` to prevent race conditions on shared data.

### **6. Test Coverage**

- Includes a comprehensive test suite under `internal/repositories/memory/tests`.

---

## Project Structure

```
.
├── cmd
│   ├── config
│   │   └── config.go           # Configuration management
│   └── server
│       └── main.go             # Service entry point
├── config.yml                   # Configuration file
├── docker-compose.yml            # Docker Compose setup
├── Dockerfile                    # Docker build instructions
├── go.mod
├── go.sum
├── internal
│   ├── entities
│   │   ├── campaign.go         # Campaign model
│   │   ├── impression.go       # Impression model
│   │   ├── server.go           # Server struct with concurrent maps
│   │   └── stats.go            # Stats model
│   ├── handlers
│   │   ├── campaign.go         # Campaign HTTP handlers
│   │   ├── impression.go       # Impression HTTP handlers
│   │   ├── notFound.go         # 404 error handler
│   │   ├── server.go           # Server initialization
│   │   └── stats.go            # Stats handler
│   ├── logger
│   │   └── logger.go           # Custom logging utilities
│   ├── repositories
│   │   ├── campaign_repository.go  # Campaign repository interface
│   │   ├── impression_repository.go # Impression repository interface
│   │   ├── memory
│   │   │   ├── campaign.go     # In-memory campaign storage
│   │   │   ├── impression.go   # In-memory impression storage
│   │   │   ├── initiate.go     # Initialization logic
│   │   │   ├── stats.go        # In-memory stats calculation
│   │   │   └── tests
│   │   │       ├── campaign_test.go
│   │   │       ├── config.yml
│   │   │       ├── high_volume_test.go
│   │   │       ├── impression_test.go
│   │   │       ├── not_found_test.go
│   │   │       ├── stats_test.go
│   │   │       └── utils.go
│   │   └── stats_repository.go # Stats repository interface
│   ├── utils
│   │   └── response.go         # API response helpers
│   └── validators
│       ├── campaign.go         # Campaign validation logic
│       ├── impression.go       # Impression validation logic
│       ├── stats.go            # Stats validation logic
│       └── validate.go         # Generic validation utilities
├── main.go                     # (Optional) Service entry point
└── README.md
```

---

## Getting Started

### **Prerequisites**

- [Go 1.18+](https://golang.org/dl/)
- (Optional) [Docker](https://docs.docker.com/get-docker/) & [Docker Compose](https://docs.docker.com/compose/install/)

### **Installation**

1. **Clone** the repository:

```bash
git clone https://github.com/yourusername/ad-impression-service.git
cd ad-impression-service
```

2. **Install dependencies**:

```bash
go mod download
```

### **Running Locally**

#### **Using Go:**

```bash
go run ./cmd/server/main.go
```

The server will start on `:8080` with output:

```
Server started on :8080
```

#### **Using Docker Compose:**

```bash
docker-compose up --build
```

The service will be available at `:8080`.

---

## Usage

### **1. Create a Campaign**

```bash
curl -X POST \
  http://localhost:8080/api/v1/campaigns \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Campaign A",
    "start_time": "2025-01-01T00:00:00Z"
}'
```

**Response:**

```json
{
  "id": "some-uuid-value",
  "name": "Campaign A",
  "start_time": "2025-01-01T00:00:00Z"
}
```

### **2. Track an Impression**

```bash
curl -X POST \
  http://localhost:8080/api/v1/impressions \
  -H 'Content-Type: application/json' \
  -d '{
    "campaign_id": "some-uuid-value",
    "user_id": "user123",
    "ad_id": "ad456"
}'
```

**Response:**

```json
{ "message": "Impression saved successfully" }
```

### **3. Get Campaign Stats**

```bash
curl -X GET http://localhost:8080/api/v1/campaigns/some-uuid-value/stats
```

**Response:**

```json
{
  "campaign_id": "some-uuid-value",
  "last_hour": 10,
  "last_day": 50,
  "total": 100
}
```

---

## Testing

Run tests:

```bash
go test -coverprofile=coverage.out ./...
```

Check coverage:

```bash
go tool cover -func=coverage.out
```

Generate an HTML coverage report:

```bash
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```

---

## Contributing

1. **Fork** the repository.
2. **Create** a feature branch: `git checkout -b feature/new-feature`
3. **Commit** your changes: `git commit -m 'Add some feature'`
4. **Push** to the branch: `git push origin feature/new-feature`
5. **Open** a Pull Request.
