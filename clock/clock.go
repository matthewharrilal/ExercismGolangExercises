package clock

// So let us understand what has to be completed here ... we have to first be able to create a new clock 

// The way the tests depict it is that they will create a new object tuple containing the hour and the minutes

// When they print out the string format of that clock instance it should output that time


type Clock struct {
	Hour int

	Minutes int
}


func (clock Clock) String {

}