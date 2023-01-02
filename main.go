package main

import (
	"code_challenge_elevator/elevator"
	"flag"
	"fmt"
	"time"
)

func main() {

	pauseTimeInSecsPtr := flag.Int("pauseTimeInSecs", 2, "Pause time in seconds between 2 states transition")
	skipPausePtr := flag.Bool("skipPause", false, "Skip the initial pause to read pictograms")
	flag.Parse()

	banner := `
 ██████╗  ██████╗      ██████╗ ██████╗ ██████╗ ███████╗     ██████╗██╗  ██╗ █████╗ ██╗     ██╗     ███████╗███╗   ██╗ ██████╗ ███████╗
██╔════╝ ██╔═══██╗    ██╔════╝██╔═══██╗██╔══██╗██╔════╝    ██╔════╝██║  ██║██╔══██╗██║     ██║     ██╔════╝████╗  ██║██╔════╝ ██╔════╝
██║  ███╗██║   ██║    ██║     ██║   ██║██║  ██║█████╗      ██║     ███████║███████║██║     ██║     █████╗  ██╔██╗ ██║██║  ███╗█████╗  
██║   ██║██║   ██║    ██║     ██║   ██║██║  ██║██╔══╝      ██║     ██╔══██║██╔══██║██║     ██║     ██╔══╝  ██║╚██╗██║██║   ██║██╔══╝  
╚██████╔╝╚██████╔╝    ╚██████╗╚██████╔╝██████╔╝███████╗    ╚██████╗██║  ██║██║  ██║███████╗███████╗███████╗██║ ╚████║╚██████╔╝███████╗
 ╚═════╝  ╚═════╝      ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝     ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝ ╚═════╝ ╚══════╝
                                                                                                                                      
███████╗██╗     ███████╗██╗   ██╗ █████╗ ████████╗ ██████╗ ██████╗                                                                    
██╔════╝██║     ██╔════╝██║   ██║██╔══██╗╚══██╔══╝██╔═══██╗██╔══██╗                                                                   
█████╗  ██║     █████╗  ██║   ██║███████║   ██║   ██║   ██║██████╔╝                                                                   
██╔══╝  ██║     ██╔══╝  ╚██╗ ██╔╝██╔══██║   ██║   ██║   ██║██╔══██╗                                                                   
███████╗███████╗███████╗ ╚████╔╝ ██║  ██║   ██║   ╚██████╔╝██║  ██║                                                                   
╚══════╝╚══════╝╚══════╝  ╚═══╝  ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═╝  ╚═╝                                                                   
                                                                                                                                      
	`
	fmt.Println(banner)

	fmt.Printf("\n\tTime between 2 states transitions set to %d seconds\n\n", *pauseTimeInSecsPtr)

	legends := `
	Pictograms:
	
	⎣x⎦ : elevator STAYING EMPTY at floor 'x'
	x☹x : people WAITING for elevator at floor 'x'
	❲x❳ : destination floor 'x'
	|⋅⟩ : EMPTY elevator moving UP
	⟨⋅| : EMPTY elevator moving DOWN
	↑x↑ : elevator LOADING people at floor 'x'
	↓x↓ : elevator UNLOADING people at floor 'x'
	|☺⟩ : elevator TRANSPORTING people moving UP
	⟨☺| : elevator TRANSPORTING people moving DOWN
	x->y: an ORDER to take people from floor 'x' to floor 'y'
	
	Display system: 
	
	1 [1->3](MovingEmptyTo)       :|⋅⟩1☹1 _ ❲3❳   means 
	
	elevator n°1, with current order Floor 1 to Floor 3, current state: MovingEmptyTo, then the display of the elevator movement 		

	Pausing 15 seconds to let you read the pictograms and understand the display system ....
`
	fmt.Print(legends)

	if !*skipPausePtr {
		time.Sleep(15 * time.Second)
	}

	controller := elevator.NewController(*pauseTimeInSecsPtr)

	controller.AddElevator(1)
	controller.AddElevator(2)

	controller.PushOrder(1, 3)
	controller.PushOrder(5, 2)
	controller.PushOrder(0, 2)
	controller.PushOrder(3, 6)
	controller.PushOrder(4, 0)

	controller.Run()

}
