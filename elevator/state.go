package elevator

import (
	"fmt"
	"strings"
)

type State interface {
	floor() Floor
	display(currentOrder Order, currentPosition int) string
}

type TransportingPeopleTo struct {
	toFloor Floor
}

func (t TransportingPeopleTo) floor() Floor {
	return t.toFloor
}

func (t TransportingPeopleTo) display(currentOrder Order, currentPosition int) string {
	from := currentOrder.from.toInt()
	to := currentOrder.to.toInt()
	display := fmt.Sprintf("%s%-22s:", currentOrder, "(TransportingPeopleTo)")

	if from < to {
		// from < to
		// currentPosition >= from && currentPosition <= to by design

		if currentPosition < to {
			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("|☺⟩")

			display += strings.Repeat(" _ ", to-currentPosition-1)

			display += fmt.Sprintf("❲%d❳", to)
		} else { // currentPosition == to
			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("%d☺%d", currentPosition, currentPosition)
		}

	} else {

		// to < from
		// to <= currentPosition && currentPosition <= from by design

		if currentPosition == to {
			display += strings.Repeat(" _ ", to)

			display += fmt.Sprintf("%d☺%d", to, to)

			display += strings.Repeat(" _ ", from-currentPosition)

		} else { // currentPosition > to

			display += strings.Repeat(" _ ", to)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", currentPosition-to-1)

			display += fmt.Sprintf("⟨☺|")

			display += strings.Repeat(" _ ", from-currentPosition)
		}
	}

	return display
}

type MovingEmptyTo struct {
	toFloor Floor
}

func (m MovingEmptyTo) floor() Floor {
	return m.toFloor
}

func (m MovingEmptyTo) display(currentOrder Order, currentPosition int) string {
	from := currentOrder.from.toInt()
	to := currentOrder.to.toInt()

	display := fmt.Sprintf("%s%-22s:", currentOrder, "(MovingEmptyTo)")

	if from < to {

		if currentPosition < from {

			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("|⋅⟩")

			display += strings.Repeat(" _ ", from-currentPosition-1)

			display += fmt.Sprintf("%d☹%d", from, from)

			display += strings.Repeat(" _ ", to-from-1)

			display += fmt.Sprintf("❲%d❳", to)

		} else if currentPosition == from {
			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("⎣%d⎦", currentPosition)

			display += strings.Repeat(" _ ", to-currentPosition-1)

			display += fmt.Sprintf("❲%d❳", to)

		} else if currentPosition > from && currentPosition < to {

			display += strings.Repeat(" _ ", from)

			display += fmt.Sprintf("%d☹%d", from, from)

			display += strings.Repeat(" _ ", currentPosition-from-1)

			display += fmt.Sprintf("⟨⋅|")

			display += strings.Repeat(" _ ", to-currentPosition-1)

			display += fmt.Sprintf("❲%d❳", to)

		} else if currentPosition == to {

			display += strings.Repeat(" _ ", from)

			display += fmt.Sprintf("%d☹%d", from, from)

			display += strings.Repeat(" _ ", currentPosition-from-1)

			display += fmt.Sprintf("⟨⋅|")

		} else if currentPosition > to {

			display += strings.Repeat(" _ ", from)

			display += fmt.Sprintf("%d☹%d", from, from)

			display += strings.Repeat(" _ ", to-from-1)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", currentPosition-to-1)

			display += fmt.Sprintf("⟨⋅|")
		}
	} else { // to < from

		if currentPosition < to {
			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("|⋅⟩")

			display += strings.Repeat(" _ ", to-currentPosition-1)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", from-to-1)

			display += fmt.Sprintf("%d☹%d", from, from)

		} else if currentPosition == to {

			display += strings.Repeat(" _ ", currentPosition)

			display += fmt.Sprintf("|⋅⟩")

			display += strings.Repeat(" _ ", from-currentPosition-1)

			display += fmt.Sprintf("%d☹%d", from, from)

		} else if currentPosition > to && currentPosition < from {
			display += strings.Repeat(" _ ", to)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", currentPosition-to-1)

			display += fmt.Sprintf("|⋅⟩")

			display += strings.Repeat(" _ ", from-currentPosition-1)

			display += fmt.Sprintf("%d☹%d", from, from)
		} else if currentPosition == from {
			display += strings.Repeat(" _ ", to)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", currentPosition-to-1)

			display += fmt.Sprintf("⎣%d⎦", currentPosition)

		} else if currentPosition > from {
			display += strings.Repeat(" _ ", to)

			display += fmt.Sprintf("❲%d❳", to)

			display += strings.Repeat(" _ ", from-to-1)

			display += fmt.Sprintf("%d☹%d", from, from)

			display += strings.Repeat(" _ ", currentPosition-from-1)

			display += fmt.Sprintf("⟨⋅|")
		}

	}

	return display
}

type StopAtFloor struct {
	currentFloor Floor
}

func (s StopAtFloor) floor() Floor {
	return s.currentFloor
}

func (s StopAtFloor) display(currentOrder Order, currentPosition int) string {
	var display string
	if (Order{}) == currentOrder {
		display = fmt.Sprintf("%s%-22s:", "[    ]", "(StopAtFloor)")

		display += strings.Repeat(" _ ", currentPosition)

		display += fmt.Sprintf("⎣%d⎦", currentPosition)

	} else {
		display = fmt.Sprintf("%s%-22s:", currentOrder, "(StopAtFloor)")

		from := currentOrder.from.toInt()
		to := currentOrder.to.toInt()

		if from < to {
			if currentPosition < from {
				display += strings.Repeat(" _ ", currentPosition)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", from-currentPosition-1)

				display += fmt.Sprintf("%d☹%d", from, from)

				display += strings.Repeat(" _ ", to-from-1)

				display += fmt.Sprintf("❲%d❳", to)

			} else if currentPosition == from {
				display += strings.Repeat(" _ ", from)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", to-from-1)

				display += fmt.Sprintf("❲%d❳", to)

			} else if currentPosition > from && currentPosition < to {
				display += strings.Repeat(" _ ", from)

				display += fmt.Sprintf("%d☹%d", from, from)

				display += strings.Repeat(" _ ", currentPosition-from-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", to-currentPosition-1)

				display += fmt.Sprintf("❲%d❳", to)
			} else if currentPosition == to {
				display += strings.Repeat(" _ ", from)

				display += fmt.Sprintf("%d☹%d", from, from)

				display += strings.Repeat(" _ ", currentPosition-from-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)
			} else if currentPosition > to {
				display += strings.Repeat(" _ ", from)

				display += fmt.Sprintf("%d☹%d", from, from)

				display += strings.Repeat(" _ ", to-from-1)

				display += fmt.Sprintf("❲%d❳", to)

				display += strings.Repeat(" _ ", currentPosition-to-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)
			}
		} else {
			// to < from

			if currentPosition < to {
				display += strings.Repeat(" _ ", currentPosition)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", to-currentPosition-1)

				display += fmt.Sprintf("❲%d❳", to)

				display += strings.Repeat(" _ ", from-to-1)

				display += fmt.Sprintf("%d☹%d", from, from)

			} else if currentPosition == to {
				display += strings.Repeat(" _ ", currentPosition)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", from-currentPosition-1)

				display += fmt.Sprintf("%d☹%d", from, from)

			} else if currentPosition > to && currentPosition < from {
				display += strings.Repeat(" _ ", to)

				display += fmt.Sprintf("❲%d❳", to)

				display += strings.Repeat(" _ ", currentPosition-to-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

				display += strings.Repeat(" _ ", from-currentPosition-1)

				display += fmt.Sprintf("%d☹%d", from, from)

			} else if currentPosition == from {
				display += strings.Repeat(" _ ", to)

				display += fmt.Sprintf("❲%d❳", to)

				display += strings.Repeat(" _ ", currentPosition-to-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)

			} else if currentPosition > from {
				display += strings.Repeat(" _ ", to)

				display += fmt.Sprintf("❲%d❳", to)

				display += strings.Repeat(" _ ", from-to-1)

				display += fmt.Sprintf("%d☹%d", from, from)

				display += strings.Repeat(" _ ", currentPosition-from-1)

				display += fmt.Sprintf("⎣%d⎦", currentPosition)
			}
		}

	}

	return display
}

type LoadingAtFloor struct {
	currentFloor Floor
}

func (l LoadingAtFloor) floor() Floor {
	return l.currentFloor
}

func (l LoadingAtFloor) display(currentOrder Order, currentPosition int) string {
	from := currentOrder.from.toInt()
	to := currentOrder.to.toInt()

	display := fmt.Sprintf("%s%-22s:", currentOrder, "(LoadingAtFloor)")

	if from < to {
		// currentPosition == from by design
		display += strings.Repeat(" _ ", from)

		display += fmt.Sprintf("↑%d↑", currentPosition)

		display += strings.Repeat(" _ ", to-currentPosition-1)

		display += fmt.Sprintf("❲%d❳", to)

	} else {
		// to < from
		// currentPosition == from by design
		display += strings.Repeat(" _ ", to)

		display += fmt.Sprintf("❲%d❳", to)

		display += strings.Repeat(" _ ", currentPosition-to-1)

		display += fmt.Sprintf("↑%d↑", currentPosition)
	}

	return display
}

type UnloadingAtFloor struct {
	currentFloor Floor
}

func (u UnloadingAtFloor) floor() Floor {
	return u.currentFloor
}

func (u UnloadingAtFloor) display(currentOrder Order, currentPosition int) string {

	display := fmt.Sprintf("%s%-22s:", currentOrder, "(UnloadingAtFloor)")

	display += strings.Repeat(" _ ", currentPosition)

	display += fmt.Sprintf("↓%d↓", currentPosition)

	return display
}
