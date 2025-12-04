import { getCurrentWeather } from "./api/weather.js";
import { createWeatherCard } from "./components/WeatherCard.js";

const input = document.getElementById("city-input");
const button = document.getElementById("search-btn");
const result = document.getElementById("weather-result");

button.addEventListener("click", async () => {
    const city = input.value.trim();
    if (!city) return;

    result.innerHTML = "<p>Loading...</p>";
    result.classList.remove("hidden");

    try {
        const data = await getCurrentWeather(city);
        if (!data.success) {
            result.innerHTML = `<p>City not found</p>`;
            return;
        }

        // Чисто одно обновление DOM
        const card = createWeatherCard(data.data);
        result.innerHTML = "";
        result.appendChild(card);
    } catch {
        result.innerHTML = `<p>Error occurred</p>`;
    }
});