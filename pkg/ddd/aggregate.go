package ddd

type (
	Aggregate struct {
		eventIDs []string
	}
)

func (a *Aggregate) AddEventID(eventID string) {
	a.eventIDs = append(a.eventIDs, eventID)
}

func (a *Aggregate) EventIDs() []string {
	return a.eventIDs
}

func NoopAggregate() Aggregate {
	return Aggregate{
		eventIDs: nil,
	}
}
