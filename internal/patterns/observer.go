package patterns

import (
	"fmt"
	"slices"
)

type Publisher interface {
	RegisterObserver(s Subscriber)
	RemoveObserver(s Subscriber)
	NotifyOberservers()
}

type Subscriber interface {
	Update(temperature float64)
}

type Weatherstation struct {
	temperature float64
	observers   []Subscriber
}

func (s *Weatherstation) SetTemperature(temperature float64) {
	s.temperature = temperature
	fmt.Println("Weatherstation: new temperature measurement: ", s.temperature)
	s.NotifyOberservers()
}

func (s *Weatherstation) RegisterObserver(observer Subscriber) {
	s.observers = append(s.observers, observer)
}

func (s *Weatherstation) RemoveObserver(observer Subscriber) {

	index := slices.Index(s.observers, observer)
	deleteElement(s.observers, index)
}

func (s *Weatherstation) NotifyOberservers() {
	for _, obs := range s.observers {
		obs.Update(s.temperature)
	}
}

/*Auxilar functions*/
func deleteElement(slice []Subscriber, index int) []Subscriber {
	return append(slice[:index], slice[index+1:]...)
}

func pop(slice []Subscriber) []Subscriber {

	if len(slice) == 0 {
		return nil
	}
	return append([]Subscriber{}, slice[:cap(slice)-1]...)
}

/*-----------------------------------------------------*/

type TemperatureDisplay struct {
	subject  Publisher
	observer Subscriber
}

func NewTemperatureDisplay(Weatherstation Publisher, tdisplay Subscriber) *TemperatureDisplay {
	w := Weatherstation
	td := tdisplay
	w.RegisterObserver(td)

	return &TemperatureDisplay{
		subject:  w,
		observer: td,
	}

}

func (td *TemperatureDisplay) Update(temperature float64) {
	fmt.Println("TemperatureDisplay: I need to update my display...")
}

type Fun struct {
	subject  Publisher
	observer Subscriber
}

/*--------------------------*/
func NewFun(Weatherstation Publisher, tfun Subscriber) *Fun {
	w := Weatherstation
	td := tfun
	w.RegisterObserver(td)

	return &Fun{
		subject:  w,
		observer: td,
	}
}

func (f *Fun) Update(temperature float64) {
	if temperature > 25 {
		fmt.Println("It's hot here, time to turn on")
	} else {
		fmt.Println("It's nice and cool, keep it off")
	}
}
