const API_BASE_URL = 'http://localhost:8000/forecast';

export async function getCurrentWeather(city) {
    try {
        const response = await fetch(`${API_BASE_URL}/current?place=${city}`);
        if (!response.ok) throw new Error("Network Error");
        return await response.json();
    }catch(err) {
        console.error("Error getting current weather", err);
        throw err;
    }
}

export async function getForecastEvery3Hours(city) {
    try {
        const response = await fetch(`${API_BASE_URL}?place=${city}`);
        if (!response.ok) throw new Error("Network Error");
        return await response.json();
    }catch(err) {
        console.error("Error getting weather", err);
        throw err;
    }
}

export async function getForecastPerDay(city) {
    try{
        const response = await fetch(`${API_BASE_URL}/per-day?place=${city}`);
        if (!response.ok) throw new Error("Network Error");
        return await response.json();
    }catch(err) {
        console.error("Error getting weather", err);
        throw err;
    }
}