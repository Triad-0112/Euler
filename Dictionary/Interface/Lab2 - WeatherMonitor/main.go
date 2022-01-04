package main

import "fmt"

type Observer interface {
	update(temp float64, humidity float64, pressure float64)
}

type Subject interface {
	registerObserver(Observer)
	deregisterObserver(Observer)
	notifyObservers()
	getTemperature()
	getHumidity()
	getPressure()
	measurementsChanged()
}

type displayElement interface {
	display()
}

//CURRENT CONDITION DISPLAY
type currentConditionDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
}

func newCurrentConditionDisplay() *currentConditionDisplay {
	return &currentConditionDisplay{}
}

func (ccd *currentConditionDisplay) update(temp float64, humidity float64, pressure float64) {
	ccd.temperature = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.display()
}
func (ccd *currentConditionDisplay) display() {
	fmt.Printf("\nCurrent conditions:\nTemperature:\t\t%.2f\nHumidity:\t\t%.2f\nPressure:\t\t%.2f\n", ccd.temperature, ccd.humidity, ccd.pressure)
}

// WEATHER DATA
type weatherData struct {
	observers               map[Observer]bool
	temperature             float64
	humidity                float64
	pressure                float64
	currentConditionDisplay currentConditionDisplay
	statisticsDisplay       statisticsDisplay
	//forecastDisplay         forecastDisplay
}

func newWeatherData() *weatherData {
	return &weatherData{
		observers: make(map[Observer]bool),
	}
}

func (w *weatherData) registerObserver(o Observer) {
	w.observers[o] = true
}

func (w *weatherData) deregisterObserver(o Observer) {
	if _, ok := w.observers[o]; ok {
		delete(w.observers, o)
	}
}

//func (w *weatherData) measurementsChanged() {
//w.notifyObservers()
//}

func (w *weatherData) setMeasurements(temp float64, humidity float64, pressure float64) {
	w.temperature = temp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}

func (w *weatherData) notifyObservers() {
	for Observer := range w.observers {
		Observer.update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *weatherData) measurementsChanged() {
	w.notifyObservers()
	temp := w.temperature
	pressure := w.pressure
	humidity := w.humidity
	w.currentConditionDisplay.update(temp, humidity, pressure)
	w.statisticsDisplay.update(temp, humidity, pressure)
	//w.forecastDisplay.update(temp, humidity, pressure)
}

//STATISTIC DISPLAY
type statisticsDisplay struct {
	count   uint32
	avgTemp float64
	maxTemp float64
	minTemp float64
}

func newStatisticsDisplay() *statisticsDisplay {
	return &statisticsDisplay{}
}

func (sd *statisticsDisplay) update(temp float64, humidity float64, pressure float64) {
	sd.count++

	sd.avgTemp -= (sd.avgTemp - temp) / float64(sd.count)

	if sd.maxTemp < temp || sd.maxTemp == 0.0 {
		sd.maxTemp = temp
	}

	if sd.minTemp > temp || sd.minTemp == 0.0 {
		sd.minTemp = temp
	}
	sd.display()
}

func (sd *statisticsDisplay) display() {
	fmt.Printf("Avg Temperature:\t%.2f\nMax Temperature:\t%.2f\nMin Temperature:\t%.2f\n\n-------\n", sd.avgTemp, sd.maxTemp, sd.minTemp)
}

func main() {
	weatherData := newWeatherData()

	currentConditionDisplay := newCurrentConditionDisplay()
	weatherData.registerObserver(currentConditionDisplay)

	statisticsDisplay := newStatisticsDisplay()
	weatherData.registerObserver(statisticsDisplay)

	weatherData.setMeasurements(80, 65, 30.4)
	weatherData.setMeasurements(82, 70, 29.2)
	weatherData.setMeasurements(77, 90, 29.2)

}
