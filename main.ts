/**
 * @author Brayden Thistle
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview this program calculates and splits the total of a meal amongst 5 people.
 */

// meal cost
let meal: number = 315.99

// amount of people splitting the bill
let splitters: number = 5

//total amount of the meal with the rates
console.log("The base cost of the meal is $" + meal + ". With a 13% HST rate the price comes up to $" + 0.13 * meal + ". With the tip rate of 15% the total cost of the meal comes to $" + (357.0687 * 0.15 + 357.0687));

// total cost of the meal
let TotalCost: number = 410.63

// the meal split amongst 5 people
console.log("since there are " + splitters + " people splitting the bill. The total comes to $" + (TotalCost / splitters) + " per person.");

console.log("\nDone.");