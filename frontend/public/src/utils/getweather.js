export function getWeatherIcon(description){
    if (description.includes("clear sky")){
        return `assets/images/day-sunny-color-icon.png`;
    }
    if (description.includes("scattered clouds") || description.includes("broken clouds")){
        return `assets/images/weather-icon.png`;
    }
    if (description.includes("overcast clouds")){
        return `assets/images/cloud-color-icon.png`;
    }
    if (description.includes("rain")){
        return `assets/images/cloud-rain-color-icon.png`;
    }
    return "";
}