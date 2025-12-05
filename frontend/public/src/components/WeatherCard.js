import {getWeatherIcon} from "../utils/getweather.js";

export function createWeatherCard(weather) {
    const weatherDesc = weather.weather.length > 0 ? weather.weather[0].description : "N/A";

    const card = document.createElement("div");
    const icon=getWeatherIcon(weatherDesc);
    card.className = "weather-card";
    card.innerHTML = `
        <div class="weather-icon"><img src="${icon}" alt="${weatherDesc}"></div>
        <div class="weather-text">
        <div class="weather-desc">${weatherDesc}</div>
        <div class="weather-temp">${weather.temp}°C (feels like ${weather.feels_like}°C)</div>
        <div class="weather-info">Min: ${weather.temp_min}°C | Max: ${weather.temp_max}°C</div>
        <div class="weather-info">Humidity: ${weather.humidity}%</div>
        <div class="weather-info">Pressure: ${weather.pressure} hPa</div>
        <div class="weather-info">Wind: ${weather.speed} m/s, ${weather.deg}°</div>
        <div class="weather-info">Clouds: ${weather.clouds}% | Visibility: ${weather.visibility} m</div>
        </div>    
    `;
    return card;
}

export function createForecastGrid(dayData) {
    const weatherDesc=dayData.weather && dayData.weather.length > 0 ? dayData.weather[0].description : "N/A";

    const date = new Date(dayData.day);
    const options = { day: "2-digit", month: "short" }; // e.g. "05 Dec"
    const dayFormatted = date.toLocaleDateString("en-US", options);

    const grid=document.createElement("div");
    grid.className = "forecast-grid";
    grid.innerHTML = `
        <div class="day-name">${dayFormatted}</div>
        <div class="temp-range">${dayData.temp_min}°C — ${dayData.temp_max}°C</div>
        <div class="weather-desc">${weatherDesc}</div>
    `;
    return grid;
}

export function createHourlyCard(hourData) {
    const weatherDesc= hourData.weather && hourData.weather.length > 0 ? hourData.weather[0].description : "N/A";
    const timeOnly=hourData.time.split(" ")[1].slice(0,5);

    const hourlyCard = document.createElement("div");
    hourlyCard.className = "hourly-card";
    hourlyCard.innerHTML = `
        <div class="time">${timeOnly}</div>
        <div class="temp">${hourData.temp}°C</div>
        <div class="weather-desc">${weatherDesc}</div>
    `;
    return hourlyCard;
}