package elevator

import (
	"errors"
	"fmt"
	"github.com/mariomac/gostream/stream"
	"golang.org/x/exp/maps"
	"reflect"
	"time"
)

type Controller struct {
	elevators       map[int]Elevator
	ordersBuffer    Orders
	pauseTimeInSecs int
}

func NewController(pauseTimeInSecs int) *Controller {
	return &Controller{
		elevators:       map[int]Elevator{},
		ordersBuffer:    Orders{},
		pauseTimeInSecs: pauseTimeInSecs,
	}
}

func (c *Controller) AddElevator(index int) bool {

	_, ok := c.elevators[index]
	if !ok {
		elevator := Elevator{
			index:        index,
			currentOrder: Order{},
			position:     Floor(0),
			state:        StopAtFloor{Floor(0)},
		}

		c.elevators[index] = elevator
		return true
	} else {
		return false
	}
}

func (c *Controller) display() string {
	display := "\n\n\tElevators state: \n"
	display += fmt.Sprintf("\tOrdersBuffer: %v\n\n", c.ordersBuffer)
	//display += "**********************************************************************\n\n"
	stream.OfSlice(maps.Values(c.elevators)).
		Sorted(sortElevatorsByIndex).
		ForEach(func(e Elevator) {
			display += fmt.Sprintf("%s\n", e.display())
		})

	return display
}

func (c *Controller) PushOrder(from int, to int) {
	newOrder := Order{from: Floor(from), to: Floor(to)}
	newBuffer := append(c.ordersBuffer, newOrder)
	c.ordersBuffer = newBuffer
}

func (c *Controller) popOrderFromBuffer() error {
	if len(c.elevators) > 0 {

		if len(c.ordersBuffer) > 0 {
			nextOrder := c.ordersBuffer[0]

			elevators := maps.Values(c.elevators)
			sortedElevators := stream.OfSlice(elevators).
				Filter(func(e Elevator) bool {
					return e.isReadyForNewOrder()
				}).
				Sorted(func(left Elevator, right Elevator) int {
					return sortElevatorsByDistance(left, right, nextOrder)
				}).
				ToSlice()

			if len(sortedElevators) > 0 {
				elevatorToUpdate := sortedElevators[0]
				newElevator, err := elevatorToUpdate.addOrder(nextOrder)
				if err == nil {
					c.elevators[elevatorToUpdate.index] = newElevator
					c.ordersBuffer = c.ordersBuffer[1:]
					return nil
				} else {
					return err
				}
			} else {
				return nil
			}

		} else {
			return nil
		}

	} else {
		return errors.New("there is no elevator configured in the system currently to receive orders")
	}
}

func sortElevatorsByDistance(left Elevator, right Elevator, newOrder Order) int {

	leftStateIsFree := reflect.TypeOf(left.state).Name() == "StopAtFloor"
	rightStateIsFree := reflect.TypeOf(right.state).Name() == "StopAtFloor"

	if leftStateIsFree && !rightStateIsFree {
		return -1
	} else if !leftStateIsFree && rightStateIsFree {
		return 1
	} else {
		leftRemainingDistance := left.remainingDistance(newOrder)
		rightRemainingDistance := right.remainingDistance(newOrder)
		if leftRemainingDistance < rightRemainingDistance {
			return -1
		} else if leftRemainingDistance > rightRemainingDistance {
			return 1
		} else {
			return 0
		}
	}
}

func sortElevatorsByIndex(left Elevator, right Elevator) int {
	if left.index < right.index {
		return -1
	} else if left.index == right.index {
		return 0
	} else {
		return 1
	}
}

func (c *Controller) Run() {

	for true {

		fmt.Println(c.display())

		newElevator := map[int]Elevator{}
		for index, v := range c.elevators {
			newElevator[index] = v.nextState()
		}
		c.elevators = newElevator

		err := c.popOrderFromBuffer()
		if err != nil {
			panic(fmt.Sprintf("%s", err))
		}

		time.Sleep(time.Duration(c.pauseTimeInSecs) * time.Second)

		if len(c.ordersBuffer) == 0 && stream.OfSlice(maps.Values(c.elevators)).AllMatch(Elevator.isReadyForNewOrder) {
			break
		}
	}

	fmt.Printf("\n\n**************** End of Simulation *******************\n\n")
}

