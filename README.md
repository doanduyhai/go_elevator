Elevator code challenge in GO
========
                                                                             
> **WARNING: this program use the recent GO generics feature, thus requiring GO version >= 1.8**

# I What is it ?
This is a coding exercise to get familiar with the GO language

The idea is to create an elevator system that dispatch orders to different elevators in the system

The complexity is to handle priorities between elevators and to code the display

The core of the animation is based on a state-machine, which transition an elevator from one state to another state

The code should have unit tests to ensure there is no regression

# II Global design

The V1 version implements the following:

1. elevator = state machine with 5 different states:
  - **`TransportingPeopleTo`**: transporting people to a target floor
  - **`MovingEmptyTo`**: moving the elevator empty to a floor to **pick people**
  - **`LoadingAtFloor`**: loading people at **source** floor 
  - **`UnloadingAtFloor`**: unloading people at **destination** floor
  - **`StopAtFloor`**: the elevator is stopped at a floor

2. the main controller has un orders buffer, at each round, an order from the buffer is dispatched to the **appropriate** elevator. How **appropriate** is an elevator is defined by a complex distance computation for an elevator from its current position to reach the **source** floor to pick people

3. There is an ASCII display system to simulate the movements of elevators. We use the following pictograms

  - ⎣x⎦ : elevator STAYING EMPTY at floor 'x'
  - x☹x : people WAITING for elevator at floor 'x'
  - ❲x❳ : destination floor 'x'
  - |⋅⟩ : EMPTY elevator moving UP
  - ⟨⋅| : EMPTY elevator moving DOWN
  - ↑x↑ : elevator LOADING people at floor 'x'
  - ↓x↓ : elevator UNLOADING people at floor 'x'
  - |☺⟩ : elevator TRANSPORTING people moving UP
  - ⟨☺| : elevator TRANSPORTING people moving DOWN
  - x->y: an ORDER to take people from floor 'x' to floor 'y'

# III Execute unit tests

Just type `go test code_challenge_elevator/elevator -v` to run all unit tests

# IV How to run

To run the program, just type `go run main.go`.

By default, there will be a screen presenting the pictograms and explaining the display system. The program will pause 
for 15 seconds to let you read and understand the display system

You can skip this annoying pause with the flag `-skipPause=true` : `go run main.go -skipPause=true`

The simulation will pause **2 seconds** between 2 states transition to display the current position and state of each elevator.

To change this pause time, you can use the flag `-pauseTimeInSecsPtr=x`: `go run main.go -pauseTimeInSecsPtr=1 -skipPause=true`




