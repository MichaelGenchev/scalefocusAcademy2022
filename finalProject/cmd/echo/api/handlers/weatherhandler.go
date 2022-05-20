package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)


func GetTemperature(ctx echo.Context) error {
    lat := ctx.Request().Header.Get("lat")
    lon := ctx.Request().Header.Get("lon")

    urlID := "https://api.openweathermap.org/data/2.5/weather?lat="+lat+"&lon="+lon+"&appid=f9344efaf6b5dbd6f7963b812c401f3a&units=metric"

    req, _ := http.NewRequest("GET", urlID, nil)
    res, _ := http.DefaultClient.Do(req)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }
    defer res.Body.Close()
    var data TemperatureAPI

    json.NewDecoder(res.Body).Decode(&data)

    result := WeatherInfo{}
    tempInt := math.Round(data.Main.Temp)
    temperature := fmt.Sprintf("%v", tempInt)
    result.FormatedTemp = temperature + "°"
    result.City = data.Name
    result.Description = data.Weather[0].Description

	fmt.Println(result)

    return ctx.JSON(http.StatusOK, result)

}

type WeatherInfo struct {
    FormatedTemp string `json:"formatedTemp"`
    Description string `json:"description"`
    City string `json:"city"`
}

type TemperatureAPI struct {
    Coord struct {
        Lon float64 `json:"lon"`
        Lat float64 `json:"lat"`
    } `json:"coord"`
    Weather []struct {
        ID          int    `json:"id"`
        Main        string `json:"main"`
        Description string `json:"description"`
        Icon        string `json:"icon"`
    } `json:"weather"`
    Base string `json:"base"`
    Main struct {
        Temp      float64 `json:"temp"`
        FeelsLike float64 `json:"feels_like"`
        TempMin   float64 `json:"temp_min"`
        TempMax   float64 `json:"temp_max"`
        Pressure  int     `json:"pressure"`
        Humidity  int     `json:"humidity"`
    } `json:"main"`
    Visibility int `json:"visibility"`
    Wind       struct {
        Speed float64 `json:"speed"`
        Deg   int     `json:"deg"`
    } `json:"wind"`
    Clouds struct {
        All int `json:"all"`
    } `json:"clouds"`
    Dt  int `json:"dt"`
    Sys struct {
        Type    int     `json:"type"`
        ID      int     `json:"id"`
        Message float64 `json:"message"`
        Country string  `json:"country"`
        Sunrise int     `json:"sunrise"`
        Sunset  int     `json:"sunset"`
    } `json:"sys"`
    Timezone int    `json:"timezone"`
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Cod      int    `json:"cod"`
}