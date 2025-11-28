# Weather App – Go Backend + JavaScript Frontend

Technologies:
* Backend
  * Go
  * Gin 
  * PostgreSQL
  * REST API
* Frontend
  * JavaScript 
  * HTML/CSS 

Deployment is done with Docker.\

A small full-stack application providing real-time weather information.

## Features

Go Backend (Gin): 
- REST API fetching and serving weather data from a third-party API.

Caching: 
- Optimized frequent requests to reduce API calls and improve performance.

JavaScript Frontend: 
 - Interactive interface communicating with the Go backend.

Error Handling & Validation: 
 - Robust request validation and clean routing structure.

Dockerized: 
 - Easy deployment using Docker containers.

Clean Architecture: 
 - Organized into service, controller, and model layers.

## Project Structure
```bash
├── backend
│   ├── cmd
│   ├── internals
│   └── pckg
├── frontend
└── README.md
```

