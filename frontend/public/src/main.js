import {getCurrentWeather, getForecastEvery3Hours, getForecastPerDay} from "./api/weather.js";
import {createForecastGrid, createHourlyCard, createWeatherCard} from "./components/WeatherCard.js";

const input = document.getElementById("city-input");
const button = document.getElementById("search-btn");
const result = document.getElementById("weather-result");

button.addEventListener("click", async () => {
    const city = input.value.trim();
    if (!city) return;

    result.innerHTML = "<p>Loading...</p>";
    result.classList.remove("hidden");

    try {
        // Get current weather
        const current = await getCurrentWeather(city);
        if (!current.success) {
            result.innerHTML = `<p>City not found</p>`;
            return;
        }

        const card = createWeatherCard(current.data);

        // Get per-day forecast
        const perDay = await getForecastPerDay(city);

        // Create wrapper
        const forecastWrapper = document.createElement("div");
        forecastWrapper.className = "forecast-wrapper";

        perDay.data.forEach(day => {
            const grid = createForecastGrid(day);
            forecastWrapper.appendChild(grid);
        });

        const perHour = await getForecastEvery3Hours(city);

        const today=new Date().toISOString().split("T")[0];
        const todayHours=perHour.data.filter(h=>h.time.startsWith(today));

        const perHourWrapper = document.createElement("div");
        perHourWrapper.className = "hourly-scroll";

        todayHours.forEach(hour => {
            const grid = createHourlyCard(hour);
            perHourWrapper.appendChild(grid);
        });


        // Render everything
        result.innerHTML = "";
        result.appendChild(card);
        result.appendChild(perHourWrapper);
        result.appendChild(forecastWrapper);

    } catch (err) {
        console.error(err);
        result.innerHTML = `<p>Error occurred</p>`;
    }
});