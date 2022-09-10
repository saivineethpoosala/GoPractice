//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// structure declaration

// type states struct{

//		Active string
//		Inactive string
//	}
type Items struct {
	Items1         string
	Items2         string
	Items3         string
	Items4         string
	TagStateItems1 string
	TagStateItems2 string
	TagStateItems3 string
	TagStateItems4 string
	// Active  bool
	// Inactive bool
}

func (P Items) ActivateTags() {

	// for _, element = range  {

	// 	if P.element != "Active" {

	// 		P.element == "Active"
	// 	}
	// }
	if P.TagStateItems1 != "ACTIVE" || P.TagStateItems2 != "ACTIVE" || P.TagStateItems3 != "ACTIVE" || P.TagStateItems4 != "ACTIVE" {
		P.TagStateItems1 = "ACTIVE"
		P.TagStateItems2 = "ACTIVE"
		P.TagStateItems3 = "ACTIVE"
		P.TagStateItems4 = "ACTIVE"
	}
	fmt.Println(" security status of " + P.Items1 + " is " + P.TagStateItems1)
	fmt.Println(" security status of " + P.Items2 + " is " + P.TagStateItems2)
	fmt.Println(" security status of " + P.Items3 + " is " + P.TagStateItems3)
	fmt.Println(" security status of " + P.Items4 + " is " + P.TagStateItems4)
}

func (Q Items) Deactivatetags() {
	if Q.TagStateItems1 == "ACTIVE" || Q.TagStateItems2 == "ACTIVE" || Q.TagStateItems3 == "ACTIVE" || Q.TagStateItems4 == "ACTIVE" {
		Q.TagStateItems1 = "INACTIVE"
		Q.TagStateItems2 = "INACTIVE"
		Q.TagStateItems3 = "INACTIVE"
		Q.TagStateItems4 = "INACTIVE"
	}
	fmt.Println(" security status of " + Q.Items1 + " is " + Q.TagStateItems1)
	fmt.Println(" security status of " + Q.Items2 + " is " + Q.TagStateItems2)
	fmt.Println(" security status of " + Q.Items3 + " is " + Q.TagStateItems3)
	fmt.Println(" security status of " + Q.Items4 + " is " + Q.TagStateItems4)
}
func checkout(slice []string) {
	/* Map := make(map[string]string)

	// Map[mySlice[0]]: Inactive
	// Map[mySlice[1]]: Inactive
	// Map[mySlice[2]]: Inactive
	// Map[mySlice[3]]: Inactive */

	// for _, element := range slice {
	if slice[4] != " " || slice[5] != "" || slice[6] != "" || slice[7] != "" {

		// if element != "" {
		// element = "INACTIVE"
		slice[4] = "INACTIVE"
		slice[5] = "INACTIVE"
		slice[6] = "INACTIVE"
		slice[7] = "INACTIVE"
	}
	fmt.Println(slice)
}

func main() {

	ItemsStatus := Items{

		Items1:         "HOME",
		Items2:         "LOCKER",
		Items3:         "CAR",
		Items4:         "BIKE",
		TagStateItems1: "ACTIVE",
		TagStateItems2: "ACTIVE",
		TagStateItems3: "ACTIVE",
		TagStateItems4: "ACTIVE",
	}
	// Map := make(map[string]string)
	mySlice := []string{"Item1", "Item2", "Item3", "Item4", "ItemStatus1", "ItemStatus2", "ItemStatus3", "ItemStatus4"}
	mySlice[0] = "HOME"
	mySlice[1] = "LOCKER"
	mySlice[2] = "CAR"
	mySlice[3] = "BIKE"
	mySlice[4] = "ACTIVE"
	mySlice[5] = "INACTIVE"
	mySlice[6] = "ACTIVE"
	mySlice[7] = "ACTIVE"
	fmt.Println(mySlice)
	ItemsStatus.ActivateTags()
	ItemsStatus.Deactivatetags()
	checkout(mySlice)
	fmt.Println(mySlice)
}
