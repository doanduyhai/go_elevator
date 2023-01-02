package elevator

import (
	"fmt"
)

type Floor int

func (f Floor) toInt() int {
	return int(f)
}

func floorFromInt(x int) Floor {
	return Floor(x)
}

type Order struct {
	from Floor
	to   Floor
}

func (o Order) String() string {
	return fmt.Sprintf("[%d->%d]", o.from.toInt(), o.to.toInt())
}

type Orders []Order

type Elevator struct {
	index        int
	currentOrder Order
	position     Floor
	state        State
}

func (e Elevator) computeDistance(from Floor, to Floor) int {
	distance := to.toInt() - from.toInt()
	if distance > 0 {
		return distance
	} else {
		return -distance
	}
}

func (e Elevator) isReadyForNewOrder() bool {
	switch e.state.(type) {
	case StopAtFloor:
		return (Order{}) == e.currentOrder
	case UnloadingAtFloor:
		return true
	default:
		return false
	}
}

func (e Elevator) remainingDistance(newOrder Order) int {
	switch e.state.(type) {
	case StopAtFloor, UnloadingAtFloor:
		return e.computeDistance(e.position, newOrder.from)
	case LoadingAtFloor, TransportingPeopleTo:
		return e.computeDistance(e.position, e.currentOrder.to) + e.computeDistance(e.currentOrder.to, newOrder.from)
	case MovingEmptyTo:
		return e.computeDistance(e.position, e.currentOrder.from) + e.computeDistance(e.currentOrder.from, e.currentOrder.to) + e.computeDistance(e.currentOrder.to, newOrder.from)
	default:
		return 0
	}

}

func (e Elevator) newPositionAndState(position int, state State) (Elevator, error) {
	if position < 0 {
		return Elevator{}, fmt.Errorf("Invalid negative position : %d", position)
	} else {
		return Elevator{
			index:        e.index,
			currentOrder: e.currentOrder,
			position:     floorFromInt(position),
			state:        state,
		}, nil
	}
}

func (e Elevator) addOrder(order Order) (Elevator, error) {
	if (Order{}) == order {
		return e, fmt.Errorf("cannot add empty order")
	} else if order.from < 0 || order.from > 9 {
		return e, fmt.Errorf("order.from %d is out of bound [0-9]", order.from.toInt())
	} else if order.to < 0 || order.to > 9 {
		return e, fmt.Errorf("order.to %d is out of bound [0-9]", order.to.toInt())
	} else if order.from == order.to {
		return e, fmt.Errorf("order.from %d should NOT be equal to order.to %d", order.from.toInt(), order.to.toInt())
	} else if e.currentOrder.to.toInt() != e.position.toInt() {
		return e, fmt.Errorf("the elevator n°%d has not reached yet its destination, cannot add new order", e.index)
	} else {
		return Elevator{
			index:        e.index,
			currentOrder: order,
			position:     e.position,
			state:        e.state,
		}, nil
	}

}

func (e Elevator) nextState() Elevator {
	var newElevator Elevator
	var newState State

	switch currentState := e.state.(type) {

	case TransportingPeopleTo:
		to := currentState.floor().toInt()
		currentPosition := e.position.toInt()

		if currentPosition == to {
			newState = UnloadingAtFloor{Floor(to)}
			newElevator, _ = e.newPositionAndState(to, newState)
		} else if currentPosition < to {
			newState = TransportingPeopleTo{Floor(to)}
			newElevator, _ = e.newPositionAndState(currentPosition+1, newState)
		} else {
			newState = TransportingPeopleTo{Floor(to)}
			newElevator, _ = e.newPositionAndState(currentPosition-1, newState)
		}
		return newElevator

	case UnloadingAtFloor:
		currentOrder := e.currentOrder
		if (Order{}) == currentOrder {
			return Elevator{
				index:        e.index,
				currentOrder: e.currentOrder,
				position:     e.position,
				state:        StopAtFloor{e.position},
			}
		} else if e.position.toInt() == currentOrder.from.toInt() {
			return Elevator{
				index:        e.index,
				currentOrder: e.currentOrder,
				position:     e.position,
				state:        LoadingAtFloor{e.position},
			}
		} else if e.position.toInt() == e.currentOrder.to.toInt() {
			return Elevator{
				index:        e.index,
				currentOrder: Order{},
				position:     e.position,
				state:        StopAtFloor{e.position},
			}
		} else {
			return Elevator{
				index:        e.index,
				currentOrder: e.currentOrder,
				position:     e.position,
				state:        MovingEmptyTo{currentOrder.from},
			}
		}

	case MovingEmptyTo:
		currentPosition := e.position.toInt()
		to := currentState.floor().toInt()

		if currentPosition == to {
			newState = LoadingAtFloor{Floor(to)}
			newElevator, _ = e.newPositionAndState(to, newState)
		} else if currentPosition < to {
			newState = currentState
			newElevator, _ = e.newPositionAndState(currentPosition+1, newState)
		} else {
			newState = currentState
			newElevator, _ = e.newPositionAndState(currentPosition-1, newState)
		}
		return newElevator

	case LoadingAtFloor:
		currentOrder := e.currentOrder
		newState = TransportingPeopleTo{currentOrder.to}
		return Elevator{
			index:        e.index,
			currentOrder: currentOrder,
			position:     e.position,
			state:        newState,
		}

	case StopAtFloor:
		currentOrder := e.currentOrder
		if (Order{}) == currentOrder {
			newState = currentState
		} else {
			if currentState.floor().toInt() == currentOrder.from.toInt() {
				newState = LoadingAtFloor{currentState.currentFloor}
			} else {
				newState = MovingEmptyTo{currentOrder.from}
			}
		}
		return Elevator{
			index:        e.index,
			currentOrder: currentOrder,
			position:     e.position,
			state:        newState,
		}

	default:
		panic(fmt.Sprintf("Unknown type: %T", currentState))
	}
}

func (e Elevator) display() string {
	stateDisplay := e.state.display(e.currentOrder, e.position.toInt())
	display := fmt.Sprintf("%d %s", e.index, stateDisplay)
	return display
}
