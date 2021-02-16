package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// marshalJSON()
	// unmarhsalJSON()
	// sorting()
	// bCrytpPackage()
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	exercise5()
}

type person struct {
	First string
	Last  string
	Age   int
}

type jSONPerson struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Custom Sorting

//For type person
type byAge []person

func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

type byFirstName []person

func (bf byFirstName) Len() int           { return len(bf) }
func (bf byFirstName) Swap(i, j int)      { bf[i], bf[j] = bf[j], bf[i] }
func (bf byFirstName) Less(i, j int) bool { return bf[i].First < bf[j].First }

// For type user
type byUserAge []user

func (a byUserAge) Len() int           { return len(a) }
func (a byUserAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byUserAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byUserLast []user

func (l byUserLast) Len() int           { return len(l) }
func (l byUserLast) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l byUserLast) Less(i, j int) bool { return l[i].Last < l[j].Last }

// Marshal this code to JSON
func exercise1() {

	type user struct {
		First string
		Age   int
	}

	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println("Format this into JSON")
	fmt.Println("Slice of byte of type users:", users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Marshaled to JSON")
	fmt.Println(string(bs))

}

//Unmarhsal this code to JSON
func exercise2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println("String literal to be converted to JSON:\n", s)
	sb := []byte(s)

	type person struct {
		First   string   `json:"First"`
		Last    string   `json:"Last"`
		Age     int      `json:"Age"`
		Sayings []string `json:"Sayings"`
	}

	var destinationByte []person

	err := json.Unmarshal(sb, &destinationByte)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nAfter conversion to JSON:\n", destinationByte)

}

//Encode & send to standard out
func exercise3() {
	type user struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Before encode:\n", users, "\n\t")
	fmt.Println("\nAfter encode:\t")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Here is the error:", err)
	}

}

// sort the []int and []string for each person
func exercise4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("More sorting!")
	fmt.Println("Unsorted:")
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println("\nSorted:")
	fmt.Println(xi)

	fmt.Println("\nUnsorted:")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println("Sorted:")
	fmt.Println(xs)
}

func exercise5() {

	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	fmt.Println("Unsorted users:")
	fmt.Println(users)

	fmt.Println("\nSorted by user age:")
	sort.Sort(byUserAge(users))
	fmt.Println(users)

	fmt.Println("\nSorted by users' last name:")
	sort.Sort(byUserLast(users))
	fmt.Println(users)

	for _, u := range users {
		fmt.Println("\n", u.First, "says:")
		for _, v := range u.Sayings {
			fmt.Println("\t", v)
		}
	}

}

func bCrytpPackage() {
	s := `password123`
	fmt.Println("Unhashed Password:", s)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hashed Password:", bs)

	//Compare plaintext password & hashed password
	fmt.Println("Comparing plaintext & hashed passwords:")
	err = bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err != nil {
		fmt.Println("Incorrect password!")
		return
	}
	fmt.Println("You are logged in!(Means the plaintext & hashed versions worked!")

}

func sorting() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"Richy", "Tony", "Sarah", "Joey", "Cassary"}
	fmt.Println("Standard Sorting")
	fmt.Println("\nUnsorted:")
	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("\nSorted:")
	fmt.Println(xi)
	fmt.Println(xs)

	//---------//
	fmt.Println("\nCustom Sort:")
	p1 := person{
		First: "Richy",
		Last:  "Condon",
		Age:   34,
	}

	p2 := person{
		First: "Sarah",
		Last:  "Lafferty",
		Age:   33,
	}

	p3 := person{
		First: "Joey",
		Last:  "Breslin",
		Age:   28,
	}

	p4 := person{
		First: "Cassary",
		Last:  "Brady",
		Age:   28,
	}

	people := []person{p1, p2, p3, p4}
	fmt.Println("Unsorted:")
	fmt.Println(people)

	sort.Sort(byAge(people))
	fmt.Println("Sorted by Age:")
	fmt.Println(people)

	sort.Sort(byFirstName(people))
	fmt.Println("Sorted by First Name:")
	fmt.Println(people)

}

func marshalJSON() {
	p1 := person{
		First: "Richy",
		Last:  "Condon",
		Age:   34,
	}

	p2 := person{
		First: "Sarah",
		Last:  "Lafferty",
		Age:   33,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	//Convert slice of s
	fmt.Println(string(bs))
}

//Takes a slice of byte & a pointer
func unmarhsalJSON() {
	//String literal declared & convert into a slice of bytes for 1 arg of Unmarhsal
	s := `[{"First":"Richy","Last":"Condon","Age":34},{"First":"Sarah","Last":"Lafferty","Age":33}]`
	biteSlice := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", biteSlice)

	//Slice of byte that JSON is marshaled too
	var people []jSONPerson

	//pass slice of bytes & pointer to the var people(&people), where JSON is unmarshaled to
	err := json.Unmarshal(biteSlice, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All Data:", people)

	fmt.Println("Ranging through:")
	for i, v := range people {
		fmt.Printf("\nPerson Number:%v\n", i)
		fmt.Println("Name:", v.First, v.Last, "\nAge:", v.Age)
	}

}
