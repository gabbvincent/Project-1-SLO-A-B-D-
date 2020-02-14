//Programmed by: Vincent Gabb
//Date: 2/12/2020
//The Pizza Party Calculator
//CSCI(20)


package main

import (
	"fmt"
	"time"
)

func main() {

//Declare variable for "people, price, and ammount"
  var people int 
  var price int 
  var ammount int

//Blank Println for space inbetween output.
  fmt.Println()

//Blank Println for space inbetween output.
  fmt.Println()

//Welcome message!
  fmt.Println("Hello! this is the Ultimate Pizza Party Calculator!")

//Delay output for 3 seconds.
  time.Sleep(10 * 150 * time.Millisecond)

//Blank Println for space inbetween output.
  fmt.Println()

//Enter the price of one pizza.
  fmt.Println("To start, Enter the price of each pizza.")

//Example of input.
  fmt.Println("Example: If the pizza was $30 enter 30.")

//Scan for "price"
  fmt.Scanln(&price)

//Blank Println for space inbetween output.
  fmt.Println()

//Enter ammount of pizzas.
  fmt.Println("Next, Enter the number of pizzas you bought.")

//Scan for "ammount"
  fmt.Scanln(&ammount)

//Blank Println for space inbetween output.
  fmt.Println()

//Enter amount of people.
  fmt.Println("Finally, Enter the ammount of people you'd like to split the price between.")

//Scan for "people"
  fmt.Scanln(&people)

//Blank Println for space inbetween output.
  fmt.Println()

//Delay output for 3 seconds.
  time.Sleep(10 * 200 * time.Millisecond)

//Add message to show computer is thinking.
  fmt.Println("Calculating...")

//Delay output for 3 seconds.
  time.Sleep(10 * 200 * time.Millisecond)

//Assign variable for "pricePerperson"
  var pricePerperson int = ammount * price / people

//Message with how much each person owes.
  fmt.Println("According to the data you've entered, each person owes","$",pricePerperson)
}
//end.