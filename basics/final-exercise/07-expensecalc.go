/*
We are building a program that helps users track their expenses. We need to create a struct called Expense to store information about an individual expense, including the name of the expense, the amount, and the date.


We need to create a method called Total that calculates the total amount spent on expenses.


Also, we need to create a method called getName on Expense struct that returns the name of the Expense.

package main

import "fmt"

// Declare the Expense struct here
// your code goes here

// Implement the Total method to calculate the total amount spent
// your code goes here

// Implement the getName method on the Expense struct here
// your code goes here

func main() {
    expenses := []Expense{
        Expense{"Grocery", 50.0, "2022-01-01"},
        Expense{"Gas", 30.0, "2022-01-02"},
        Expense{"Restaurant", 40.0, "2022-01-03"},
    }

    fmt.Println(Total(expenses))
    fmt.Println(expenses[0].getName())
}

Expected Output:
120
Grocery
*/