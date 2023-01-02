package elevator

import (
	"reflect"
	"strings"
	"testing"
)

func TestElevator_computeDistance(t *testing.T) {
	elevator := Elevator{
		index:        1,
		currentOrder: Order{},
		position:     Floor(0),
	}

	tests := []struct {
		name     string
		elevator Elevator
		args     Order
		want     int
	}{
		{
			name:     "nominal",
			elevator: elevator,
			args:     Order{from: Floor(2), to: Floor(5)},
			want:     3,
		},
		{
			name:     "negative",
			elevator: elevator,
			args:     Order{from: Floor(5), to: Floor(1)},
			want:     4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := elevator.computeDistance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("computeDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElevator_newPositionAndState(t *testing.T) {
	elevator := Elevator{
		index:        1,
		currentOrder: Order{},
		position:     Floor(0),
		state:        StopAtFloor{Floor(0)},
	}

	tests := []struct {
		name        string
		elevator    Elevator
		newPosition int
		newState    State
		want        Elevator
	}{
		{
			name:        "nominal",
			elevator:    elevator,
			newPosition: 3,
			newState:    LoadingAtFloor{Floor(3)},
			want: Elevator{
				index:        1,
				currentOrder: elevator.currentOrder,
				position:     Floor(3),
				state:        LoadingAtFloor{Floor(3)},
			},
		},
		{
			name:        "negative position",
			elevator:    elevator,
			newPosition: -3,
			newState:    StopAtFloor{Floor(0)},
			want: Elevator{
				index:        1,
				currentOrder: elevator.currentOrder,
				position:     Floor(0),
				state:        StopAtFloor{Floor(0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := elevator
			got, err := e.newPositionAndState(tt.newPosition, tt.newState)
			if tt.newPosition < 0 {
				if err == nil {
					t.Errorf("newPositionAndState() = %v, want %v", got, tt.want)
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPositionAndState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElevator_remainingDistance(t *testing.T) {
	tests := []struct {
		name     string
		elevator Elevator
		newOrder Order
		want     int
	}{
		{
			name: "transporting-at-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(3)},
				position:     2,
				state:        TransportingPeopleTo{Floor(3)},
			},
			newOrder: Order{from: Floor(4), to: Floor(2)},
			want:     2,
		},
		{
			name: "stopped-at-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     2,
				state:        StopAtFloor{Floor(2)},
			},
			newOrder: Order{from: Floor(5), to: Floor(2)},
			want:     3,
		},
		{
			name: "unloading-at-target-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(3)},
				position:     Floor(3),
				state:        UnloadingAtFloor{Floor(3)},
			},
			newOrder: Order{from: Floor(4), to: Floor(2)},
			want:     1,
		},
		{
			name: "descending",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(1)},
				position:     Floor(4),
				state:        TransportingPeopleTo{Floor(1)},
			},
			newOrder: Order{from: Floor(2), to: Floor(3)},
			want:     4,
		},
		{
			name: "loading-at-source-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(3)},
				position:     Floor(5),
				state:        LoadingAtFloor{Floor(5)},
			},
			newOrder: Order{from: Floor(6), to: Floor(2)},
			want:     5,
		},
		{
			name: "moving-empty-scenario",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(3)},
				position:     Floor(1),
				state:        MovingEmptyTo{Floor(5)},
			},
			newOrder: Order{from: Floor(4), to: Floor(2)},
			want:     7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.elevator.remainingDistance(tt.newOrder); got != tt.want {
				t.Errorf("remainingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor_toInt(t *testing.T) {
	tests := []struct {
		name string
		f    Floor
		want int
	}{
		{
			name: "nominal",
			f:    Floor(3),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.toInt(); got != tt.want {
				t.Errorf("toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floorFromInt(t *testing.T) {
	tests := []struct {
		name string
		args int
		want Floor
	}{
		{
			name: "nominal",
			args: 3,
			want: Floor(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floorFromInt(tt.args); got != tt.want {
				t.Errorf("floorFromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElevator_addOrder(t *testing.T) {
	tests := []struct {
		name     string
		elevator Elevator
		newOrder Order
		want     Elevator
	}{
		{
			name: "nominal",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     Floor(0),
				state:        StopAtFloor{Floor(0)},
			},
			newOrder: Order{from: Floor(6), to: Floor(5)},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(6), to: Floor(5)},
				position:     Floor(0),
				state:        StopAtFloor{Floor(0)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got, _ := tt.elevator.addOrder(tt.newOrder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOrder() = \n%+v\n, want \n%+v\n", got, tt.want)
			}
		})
	}
}

func TestElevator_addOrder_failures(t *testing.T) {
	tests := []struct {
		name       string
		elevator   Elevator
		newOrder   Order
		failureMsg string
	}{
		{
			name: "order.from-negative",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     1,
			},
			newOrder:   Order{from: Floor(-1), to: Floor(2)},
			failureMsg: "order.from -1 is out of bound [0-9]",
		},
		{
			name: "order.from-too-big",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     1,
			},
			newOrder:   Order{from: Floor(10), to: Floor(2)},
			failureMsg: "order.from 10 is out of bound [0-9]",
		},
		{
			name: "order.to-negative",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     1,
			},
			newOrder:   Order{from: Floor(1), to: Floor(-2)},
			failureMsg: "order.to -2 is out of bound [0-9]",
		},
		{
			name: "order.to-too-big",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     1,
			},
			newOrder:   Order{from: Floor(1), to: Floor(10)},
			failureMsg: "order.to 10 is out of bound [0-9]",
		},
		{
			name: "order.from-equals-order.to",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     1,
			},
			newOrder:   Order{from: Floor(1), to: Floor(1)},
			failureMsg: "order.from 1 should NOT be equal to order.to 1",
		},
		{
			name: "not-reached-destination",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(3)},
				position:     2,
			},
			newOrder:   Order{from: Floor(1), to: Floor(5)},
			failureMsg: "the elevator n°1 has not reached yet its destination, cannot add new order",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if _, err := tt.elevator.addOrder(tt.newOrder); err.Error() != tt.failureMsg {
				t.Errorf("addOrder() failure message = \n%+v\n, expected = \n%+v\n", err.Error(), tt.failureMsg)
			}
		})
	}

}

type UnknownState struct {
}

func (u UnknownState) floor() Floor {
	return Floor(0)
}

func (u UnknownState) display(currentOrder Order, currentPosition int) string {
	return ""
}

func TestElevator_nextState(t *testing.T) {

	tests := []struct {
		name         string
		currentState Elevator
		want         Elevator
	}{
		{
			name: "transporting-to-ascending",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(5)},
				position:     2,
				state:        TransportingPeopleTo{Floor(5)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(5)},
				position:     3,
				state:        TransportingPeopleTo{Floor(5)},
			},
		},
		{
			name: "transporting-to-descending",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(1)},
				position:     3,
				state:        TransportingPeopleTo{Floor(1)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(1)},
				position:     2,
				state:        TransportingPeopleTo{Floor(1)},
			},
		},
		{
			name: "transporting-to-descending",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(1)},
				position:     3,
				state:        TransportingPeopleTo{Floor(1)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(1)},
				position:     2,
				state:        TransportingPeopleTo{Floor(1)},
			},
		},
		{
			name: "transporting-ascending-unloading",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(4)},
				position:     4,
				state:        TransportingPeopleTo{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(4)},
				position:     4,
				state:        UnloadingAtFloor{Floor(4)},
			},
		},
		{
			name: "transporting-descending-unloading",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     1,
				state:        TransportingPeopleTo{Floor(1)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     1,
				state:        UnloadingAtFloor{Floor(1)},
			},
		},
		{
			name: "unloading-at-floor-and-stop",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     5,
				state:        UnloadingAtFloor{Floor(5)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     5,
				state:        StopAtFloor{Floor(5)},
			},
		},
		{
			name: "unloading-at-floor-and-loading",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(3)},
				position:     5,
				state:        UnloadingAtFloor{Floor(5)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(5), to: Floor(3)},
				position:     5,
				state:        LoadingAtFloor{Floor(5)},
			},
		},
		{
			name: "unloading-at-floor-and-moving-to",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(3)},
				position:     5,
				state:        UnloadingAtFloor{Floor(5)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(3)},
				position:     5,
				state:        MovingEmptyTo{Floor(4)},
			},
		},
		{
			name: "moving-empty-ascending",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(3), to: Floor(5)},
				position:     1,
				state:        MovingEmptyTo{Floor(3)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(3), to: Floor(5)},
				position:     2,
				state:        MovingEmptyTo{Floor(3)},
			},
		},
		{
			name: "moving-empty-descending",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     5,
				state:        MovingEmptyTo{Floor(1)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     4,
				state:        MovingEmptyTo{Floor(1)},
			},
		},
		{
			name: "moving-empty-ascending-loading",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(3), to: Floor(5)},
				position:     3,
				state:        MovingEmptyTo{Floor(3)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(3), to: Floor(5)},
				position:     3,
				state:        LoadingAtFloor{Floor(3)},
			},
		},
		{
			name: "moving-empty-descending-loading",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     1,
				state:        MovingEmptyTo{Floor(1)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     1,
				state:        LoadingAtFloor{Floor(1)},
			},
		},
		{
			name: "loading-at-floor",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     4,
				state:        LoadingAtFloor{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     4,
				state:        TransportingPeopleTo{Floor(1)},
			},
		},
		{
			name: "stop-at-floor-remaining-stopped-if-no-new-order",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
		},
		{
			name: "stop-at-floor-remaining-stopped-waiting-for-new-order",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
		},
		{
			name: "stop-at-floor-loading-from-new-order",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(4), to: Floor(1)},
				position:     4,
				state:        LoadingAtFloor{Floor(4)},
			},
		},
		{
			name: "stop-at-floor-moving-to-floor-from-new-order",
			currentState: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(2), to: Floor(5)},
				position:     4,
				state:        StopAtFloor{Floor(4)},
			},
			want: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(2), to: Floor(5)},
				position:     4,
				state:        MovingEmptyTo{Floor(2)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.currentState.nextState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextState() = \n%+v\n, want \n%+v\n", got, tt.want)
			}
		})
	}
}

func TestElevator_nextState_panic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code should have paniced")
		} else if !strings.Contains(r.(string), "Unknown type") {
			t.Errorf("Error message : %s", r)
		}
	}()

	currentState := Elevator{
		index:        1,
		currentOrder: Order{},
		position:     1,
		state:        UnknownState{},
	}

	currentState.nextState()
}

func TestElevator_display(t *testing.T) {
	tests := []struct {
		name     string
		elevator Elevator
		want     string
	}{
		{
			name: "nominal",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(2), to: Floor(4)},
				position:     4,
				state:        UnloadingAtFloor{4},
			},
			want: "1 [2->4](UnloadingAtFloor)    : _  _  _  _ ↓4↓",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.elevator.display(); got != tt.want {
				t.Errorf("String() = \n%v\n, want \n%v\n", got, tt.want)
			}
		})
	}
}

func TestElevator_isReadyForNewOrder(t *testing.T) {

	tests := []struct {
		name     string
		elevator Elevator
		want     bool
	}{
		{
			name: "stopped-at-floor-no-order",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{},
				position:     3,
				state:        StopAtFloor{Floor(3)},
			},
			want: true,
		},
		{
			name: "stopped-at-floor-but-one-order",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(4)},
				position:     3,
				state:        StopAtFloor{Floor(3)},
			},
			want: false,
		},
		{
			name: "unloading-at-target-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(4)},
				position:     4,
				state:        UnloadingAtFloor{Floor(4)},
			},
			want: true,
		},
		{
			name: "transporting-to-floor",
			elevator: Elevator{
				index:        1,
				currentOrder: Order{from: Floor(1), to: Floor(4)},
				position:     2,
				state:        TransportingPeopleTo{Floor(3)},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.elevator.isReadyForNewOrder(); got != tt.want {
				t.Errorf("isReadyForNewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
