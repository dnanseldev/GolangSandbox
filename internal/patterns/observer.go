package patterns

import "slices"

type Publisher interface {
	RegisterObserver(s Subscriber)
	RemoveObserver(s Subscriber)
	NotifyOberservers()
}

type Subscriber interface {
	Update(temperature float64)
}

type Subject struct {
	temperature float64
	observers   []Subscriber
}

func (s *Subject) RegisterObserver(observer Subscriber) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) RemoveObserver(observer Subscriber) {

	index := slices.Index(s.observers, observer)
	deleteElement(s.observers, index)
}

func (s *Subject) NotifyOberservers() {
	for _, obs := range s.observers {
		obs.Update(s.temperature)
	}
}

func (s *Subject) SetTemperature(temperature float64) {
	s.temperature = temperature
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
