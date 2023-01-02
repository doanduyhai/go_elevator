package elevator

import "testing"

func TestTransportingPeopleTo_display(t1 *testing.T) {

	tests := []struct {
		name            string
		currentOrder    Order
		currentPosition int
		currentState    State
		want            string
	}{
		// TransportingPeopleTo
		{
			name:            "transporting-ascending-in-middle",
			currentOrder:    Order{from: Floor(1), to: Floor(4)},
			currentPosition: 2,
			currentState:    TransportingPeopleTo{Floor(4)},
			want:            "[1->4](TransportingPeopleTo): _  _ |☺⟩ _ ❲4❳",
		},
		{
			name:            "transporting-ascending-at-source-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(4)},
			currentPosition: 1,
			currentState:    TransportingPeopleTo{Floor(4)},
			want:            "[1->4](TransportingPeopleTo): _ |☺⟩ _  _ ❲4❳",
		},
		{
			name:            "transporting-ascending-at-ground-floor",
			currentOrder:    Order{from: Floor(0), to: Floor(4)},
			currentPosition: 0,
			currentState:    TransportingPeopleTo{Floor(4)},
			want:            "[0->4](TransportingPeopleTo):|☺⟩ _  _  _ ❲4❳",
		},
		{
			name:            "transporting-ascending-arriving-at-target",
			currentOrder:    Order{from: Floor(0), to: Floor(4)},
			currentPosition: 4,
			currentState:    TransportingPeopleTo{Floor(4)},
			want:            "[0->4](TransportingPeopleTo): _  _  _  _ 4☺4",
		},
		{
			name:            "transporting-descending-in-middle",
			currentOrder:    Order{from: Floor(4), to: Floor(1)},
			currentPosition: 3,
			currentState:    TransportingPeopleTo{Floor(1)},
			want:            "[4->1](TransportingPeopleTo): _ ❲1❳ _ ⟨☺| _ ",
		},
		{
			name:            "transporting-descending-at-source-floor",
			currentOrder:    Order{from: Floor(4), to: Floor(1)},
			currentPosition: 4,
			currentState:    TransportingPeopleTo{Floor(1)},
			want:            "[4->1](TransportingPeopleTo): _ ❲1❳ _  _ ⟨☺|",
		},
		{
			name:            "transporting-descending-before-destination-floor",
			currentOrder:    Order{from: Floor(4), to: Floor(1)},
			currentPosition: 2,
			currentState:    TransportingPeopleTo{Floor(1)},
			want:            "[4->1](TransportingPeopleTo): _ ❲1❳⟨☺| _  _ ",
		},
		{
			name:            "transporting-descending-at-destination-floor",
			currentOrder:    Order{from: Floor(4), to: Floor(1)},
			currentPosition: 1,
			currentState:    TransportingPeopleTo{Floor(1)},
			want:            "[4->1](TransportingPeopleTo): _ 1☺1 _  _  _ ",
		},

		// MovingEmptyTo
		{
			name:            "moving-empty-ascending-before-start-floor",
			currentOrder:    Order{from: Floor(2), to: Floor(4)},
			currentPosition: 0,
			currentState:    MovingEmptyTo{Floor(3)},
			want:            "[2->4](MovingEmptyTo)       :|⋅⟩ _ 2☹2 _ ❲4❳",
		},
		{
			name:            "moving-empty-ascending-after-start-floor-before-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(5)},
			currentPosition: 3,
			currentState:    MovingEmptyTo{Floor(1)},
			want:            "[1->5](MovingEmptyTo)       : _ 1☹1 _ ⟨⋅| _ ❲5❳",
		},
		{
			name:            "moving-empty-ascending-at-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(5)},
			currentPosition: 5,
			currentState:    MovingEmptyTo{Floor(1)},
			want:            "[1->5](MovingEmptyTo)       : _ 1☹1 _  _  _ ⟨⋅|",
		},
		{
			name:            "moving-empty-ascending-after-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(5)},
			currentPosition: 6,
			currentState:    MovingEmptyTo{Floor(1)},
			want:            "[1->5](MovingEmptyTo)       : _ 1☹1 _  _  _ ❲5❳⟨⋅|",
		},
		{
			name:            "moving-empty-descending-before-target-floor",
			currentOrder:    Order{from: Floor(5), to: Floor(2)},
			currentPosition: 0,
			currentState:    MovingEmptyTo{Floor(5)},
			want:            "[5->2](MovingEmptyTo)       :|⋅⟩ _ ❲2❳ _  _ 5☹5",
		},
		{
			name:            "moving-empty-descending-at-target-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(2)},
			currentPosition: 2,
			currentState:    MovingEmptyTo{Floor(6)},
			want:            "[6->2](MovingEmptyTo)       : _  _ |⋅⟩ _  _  _ 6☹6",
		},
		{
			name:            "moving-empty-descending-between-target-and-start-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(2)},
			currentPosition: 4,
			currentState:    MovingEmptyTo{Floor(6)},
			want:            "[6->2](MovingEmptyTo)       : _  _ ❲2❳ _ |⋅⟩ _ 6☹6",
		},
		{
			name:            "moving-empty-descending-after-start-floor",
			currentOrder:    Order{from: Floor(4), to: Floor(2)},
			currentPosition: 6,
			currentState:    MovingEmptyTo{Floor(4)},
			want:            "[4->2](MovingEmptyTo)       : _  _ ❲2❳ _ 4☹4 _ ⟨⋅|",
		},

		//StopAtFloor
		{
			name:            "stop-at-floor-no-order",
			currentOrder:    Order{},
			currentPosition: 3,
			currentState:    StopAtFloor{Floor(3)},
			want:            "[    ](StopAtFloor)         : _  _  _ ⎣3⎦",
		},
		{
			name:            "stop-at-floor-ascending-before-start-floor",
			currentOrder:    Order{from: Floor(3), to: Floor(6)},
			currentPosition: 1,
			currentState:    StopAtFloor{Floor(1)},
			want:            "[3->6](StopAtFloor)         : _ ⎣1⎦ _ 3☹3 _  _ ❲6❳",
		},
		{
			name:            "stop-at-floor-ascending-at-start-floor",
			currentOrder:    Order{from: Floor(3), to: Floor(6)},
			currentPosition: 3,
			currentState:    StopAtFloor{Floor(3)},
			want:            "[3->6](StopAtFloor)         : _  _  _ ⎣3⎦ _  _ ❲6❳",
		},
		{
			name:            "stop-at-floor-ascending-after-start-floor-before-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(6)},
			currentPosition: 3,
			currentState:    StopAtFloor{Floor(3)},
			want:            "[1->6](StopAtFloor)         : _ 1☹1 _ ⎣3⎦ _  _ ❲6❳",
		},
		{
			name:            "stop-at-floor-ascending-at-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(6)},
			currentPosition: 6,
			currentState:    StopAtFloor{Floor(6)},
			want:            "[1->6](StopAtFloor)         : _ 1☹1 _  _  _  _ ⎣6⎦",
		},
		{
			name:            "stop-at-floor-ascending-after-target-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(6)},
			currentPosition: 8,
			currentState:    StopAtFloor{Floor(6)},
			want:            "[1->6](StopAtFloor)         : _ 1☹1 _  _  _  _ ❲6❳ _ ⎣8⎦",
		},
		{
			name:            "stop-at-floor-descending-before-target-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(3)},
			currentPosition: 1,
			currentState:    StopAtFloor{Floor(1)},
			want:            "[6->3](StopAtFloor)         : _ ⎣1⎦ _ ❲3❳ _  _ 6☹6",
		},
		{
			name:            "stop-at-floor-descending-at-target-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(3)},
			currentPosition: 3,
			currentState:    StopAtFloor{Floor(3)},
			want:            "[6->3](StopAtFloor)         : _  _  _ ⎣3⎦ _  _ 6☹6",
		},
		{
			name:            "stop-at-floor-descending-after-target-floor-before-start-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(1)},
			currentPosition: 3,
			currentState:    StopAtFloor{Floor(3)},
			want:            "[6->1](StopAtFloor)         : _ ❲1❳ _ ⎣3⎦ _  _ 6☹6",
		},
		{
			name:            "stop-at-floor-descending-at-start-floor",
			currentOrder:    Order{from: Floor(6), to: Floor(1)},
			currentPosition: 6,
			currentState:    StopAtFloor{Floor(6)},
			want:            "[6->1](StopAtFloor)         : _ ❲1❳ _  _  _  _ ⎣6⎦",
		},
		{
			name:            "stop-at-floor-descending-after-start-floor",
			currentOrder:    Order{from: Floor(3), to: Floor(1)},
			currentPosition: 5,
			currentState:    StopAtFloor{Floor(5)},
			want:            "[3->1](StopAtFloor)         : _ ❲1❳ _ 3☹3 _ ⎣5⎦",
		},

		//LoadingAtFloor
		{
			name:            "loading-at-floor-ascending",
			currentOrder:    Order{from: Floor(1), to: Floor(4)},
			currentPosition: 1,
			currentState:    LoadingAtFloor{Floor(1)},
			want:            "[1->4](LoadingAtFloor)      : _ ↑1↑ _  _ ❲4❳",
		},
		{
			name:            "loading-at-floor-descending",
			currentOrder:    Order{from: Floor(4), to: Floor(1)},
			currentPosition: 4,
			currentState:    LoadingAtFloor{Floor(4)},
			want:            "[4->1](LoadingAtFloor)      : _ ❲1❳ _  _ ↑4↑",
		},

		//UnloadingAtFloor
		{
			name:            "unloading-at-floor",
			currentOrder:    Order{from: Floor(1), to: Floor(4)},
			currentPosition: 4,
			currentState:    UnloadingAtFloor{Floor(4)},
			want:            "[1->4](UnloadingAtFloor)    : _  _  _  _ ↓4↓",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {

			if got := tt.currentState.display(tt.currentOrder, tt.currentPosition); got != tt.want {
				t1.Errorf("display() = \n%v\n, but wanted = \n%v\n", got, tt.want)
			}
		})
	}
}
