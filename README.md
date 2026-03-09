# WEZZAP

Wezzap is a modern web application for displaying real-time weather. The project uses Go (Gin) on the backend and pure JavaScript on the frontend, with Docker for convenient deployment.
The application receives weather data via the OpenWeather API and displays it in a user-friendly interface.


##  Project Structure
```bash
├── backend              # Go backend
│   ├── cmd/main.go      # Entry point
│   ├── configs/config.yml # Configurations
│   ├── Dockerfile
│   └── internal         # Internal modules
│       ├── client/weather.go  # Working with the OpenWeather API
│       ├── handler            # HTTP handlers
│       ├── models             # Data structures
│       ├── server/server.go   # Server configuration
│       └── services           # Data processing logic
├── frontend             # JavaScript frontend
│   ├── Dockerfile
│   ├── package.json
│   └── public
│       ├── assets
│       ├── index.html
│       └── src
│           ├── api/weather.js
│           ├── components/WeatherCard.js
│           └── utils/getweather.js
├── docker-compose.yml
└── README.md
```

## 💻 Technologies:
- Backend: Go + Gin, REST API, Graceful shutdown, clean architecture
- Frontend: JavaScript, HTML, CSS, responsive interface
- Docker: Containerisation for frontend and backend
- API: OpenWeather for weather data retrieval

## 🚀 Installation and launch
### 1. Clone the repository
```bash
git clone https://github.com/illegitie/Wezzap
cd Wezzap
```
### 2. Launching via Docker Compose
```bash
docker-compose up --build
```
- The backend will be available at http://localhost:8000
- The frontend will be available at http://localhost:8080 (or 80, if you change the port)
### 3. Configuration settings
- Backend configuration file: backend/configs/config.yml
- You can set the OpenWeather key and other parameters.
## 📌 Example of an API request
```bash
GET http://localhost:8000/forecast?place=London
```
## 📰 Response
```json
{
 {"success":true,
 "data":
  {
   "day":"2025-12-07 18:22:55",
  "temp":14.1,
  "feels_like":13.86,
  "temp_min":13.36,
  "temp_max":14.86,
  "pressure":998,
  "humidity":88,
  "speed":7.72,
  "deg":220,
  "clouds":75,
  "visibility":10000,
  "weather":
    [{
      "id":803,
      "main":"Clouds",
      "description":"broken clouds",
      "icon":"04n"
     }],
     "name":"London"
    } 
 }
}
```

## 👨‍💻 Code style
- The backend is organised according to clean architecture principles.
- The frontend is divided into components and utilities for ease of extension.
