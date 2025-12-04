export function createWeatherCard(weather) {
    const weatherDesc = weather.weather.length > 0 ? weather.weather[0].description : "N/A";

    const card = document.createElement("div");
    card.className = "weather-card";
    card.innerHTML = `
        <div class="weather-city">${weather.name}</div>
        <div class="weather-temp">${weather.temp}°C (feels like ${weather.feels_like}°C)</div>
        <div class="weather-desc">${weatherDesc}</div>
        <div>Min: ${weather.temp_min}°C | Max: ${weather.temp_max}°C</div>
        <div>Humidity: ${weather.humidity}%</div>
        <div>Pressure: ${weather.pressure} hPa</div>
        <div>Wind: ${weather.speed} m/s, ${weather.deg}°</div>
        <div>Clouds: ${weather.clouds}% | Visibility: ${weather.visibility} m</div>
    `;
    return card;
}

