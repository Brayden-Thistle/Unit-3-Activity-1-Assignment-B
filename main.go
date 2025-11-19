/*
* @author Brayden Thistle
* @version 1.0.0
* @date 2025-11-18
* @fileoverfiew this program calculates and splits the total of a meal amongst 5 people.

*/

package main

import "fmt"

func main () {
//creating variables for the cost of the meal and the amount of people splitting the bill
var meal float32 = 315.99
var splitters float64 = 5

// calculating the total cost
fmt.Println ("The base cost of the meal is $315.99. With a 13% HST rate the price comes up to $", 0.13 * meal, ". With the tip rate of 15% the total cost of the meal comes to $", (357.0687 * 0.15 + 357.0687), ".")

// creating a variable for the total cost of the meal with the HST and 
var TotalCost float64  = 410.63

//calculating the amount each person is paying.
fmt.Println ("since there are ", splitters, "people splitting the bill, each person will have to pay $", (TotalCost / splitters), "for the meal.")
}