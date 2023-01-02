package elevator

import (
	"github.com/mariomac/gostream/stream"
	"reflect"
	"testing"
)

func TestController_PushOrder(t *testing.T) {

	tests := []struct {
		name       string
		controller *Controller
		newOrder   Order
		want       Orders
	}{
		{
			name: "nominal",
			controller: &Controller{
				elevators:    map[int]Elevator{},
				ordersBuffer: Orders{Order{from: Floor(1), to: Floor(3)}},
			},
			newOrder: Order{from: Floor(4), to: Floor(2)},
			want:     Orders{Order{from: Floor(1), to: Floor(3)}, Order{from: Floor(4), to: Floor(2)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.controller.PushOrder(tt.newOrder.from.toInt(), tt.newOrder.to.toInt())

			if !reflect.DeepEqual(tt.controller.ordersBuffer, tt.want) {
				t.Errorf("curernt ordersBuffer = \n%+v\n, wanted \n%+v\n", tt.controller.ordersBuffer, tt.want)
			}
		})
	}
}

func Test_sortElevatorsByAvailability(t *testing.T) {

	tests := []struct {
		name      string
		elevators []Elevator
		newOrder  Order
		want      []Elevator
	}{
		{
			name: "one-elevator-stopped-at-floor",
			elevators: []Elevator{
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
				{
					index:        1,
					currentOrder: Order{},
					position:     5,
					state:        StopAtFloor{Floor(5)},
				},
			},
			newOrder: Order{from: Floor(1), to: Floor(3)},
			want: []Elevator{
				{
					index:        1,
					currentOrder: Order{},
					position:     5,
					state:        StopAtFloor{Floor(5)},
				},
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
			},
		},
		{
			name: "one-elevator-unloading",
			elevators: []Elevator{
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
				{
					index:        1,
					currentOrder: Order{},
					position:     2,
					state:        UnloadingAtFloor{Floor(2)},
				},
			},
			newOrder: Order{from: Floor(1), to: Floor(3)},
			want: []Elevator{
				{
					index:        1,
					currentOrder: Order{},
					position:     2,
					state:        UnloadingAtFloor{Floor(2)},
				},
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
			},
		},
		{
			name: "one-elevator-loading-at-floor",
			elevators: []Elevator{
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
				{
					index:        1,
					currentOrder: Order{from: Floor(2), to: Floor(3)},
					position:     2,
					state:        LoadingAtFloor{Floor(2)},
				},
			},
			newOrder: Order{from: Floor(1), to: Floor(3)},
			want: []Elevator{
				{
					index:        1,
					currentOrder: Order{from: Floor(2), to: Floor(3)},
					position:     2,
					state:        LoadingAtFloor{Floor(2)},
				},
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
			},
		},
		{
			name: "one-elevator-moving-empty-to",
			elevators: []Elevator{
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
				{
					index:        1,
					currentOrder: Order{from: Floor(2), to: Floor(5)},
					position:     1,
					state:        MovingEmptyTo{Floor(2)},
				},
			},
			newOrder: Order{from: Floor(1), to: Floor(3)},
			want: []Elevator{
				{
					index:        2,
					currentOrder: Order{from: Floor(1), to: Floor(4)},
					position:     2,
					state:        TransportingPeopleTo{Floor(4)},
				},
				{
					index:        1,
					currentOrder: Order{from: Floor(2), to: Floor(5)},
					position:     1,
					state:        MovingEmptyTo{Floor(2)},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortedElevators := stream.OfSlice(tt.elevators).
				Sorted(func(left Elevator, right Elevator) int {
					return sortElevatorsByDistance(left, right, tt.newOrder)
				}).
				ToSlice()
			if !reflect.DeepEqual(sortedElevators, tt.want) {
				t.Errorf("sortElevatorsByDistance() = \n%+v\n, want \n%+v\n", sortedElevators, tt.want)
			}
		})
	}
}

func TestController_popOrderFromBuffer(t *testing.T) {
	tests := []struct {
		name       string
		controller Controller
		want       Controller
	}{
		{
			name: "pop-order-from-buffer",
			controller: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(1), to: Floor(3)},
						position:     3,
						state:        UnloadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{Order{from: Floor(1), to: Floor(5)}},
			},
			want: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(1), to: Floor(5)},
						position:     3,
						state:        UnloadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{},
			},
		},
		{
			name: "no-order-from-buffer",
			controller: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(1), to: Floor(3)},
						position:     3,
						state:        UnloadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{},
			},
			want: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(1), to: Floor(3)},
						position:     3,
						state:        UnloadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{},
			},
		},
		{
			name: "no-available-elevator",
			controller: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(3), to: Floor(1)},
						position:     3,
						state:        LoadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{Order{from: Floor(1), to: Floor(6)}},
			},
			want: Controller{
				elevators: map[int]Elevator{
					1: {
						index:        1,
						currentOrder: Order{from: Floor(3), to: Floor(1)},
						position:     3,
						state:        LoadingAtFloor{Floor(3)},
					},

					2: {
						index:        2,
						currentOrder: Order{from: Floor(4), to: Floor(2)},
						position:     3,
						state:        TransportingPeopleTo{Floor(2)},
					},
				},
				ordersBuffer: Orders{Order{from: Floor(1), to: Floor(6)}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			(&tt.controller).popOrderFromBuffer()

			if !reflect.DeepEqual(tt.controller, tt.want) {
				t.Errorf("popOrderFromBuffer() actual = \n%+v\n, want \n%+v\n", tt.controller, tt.want)
			}
		})
	}
}
