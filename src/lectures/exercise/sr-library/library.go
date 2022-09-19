//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import "fmt"

type book struct {
	name    string
	checkedout bool
    checkoutby string
}

type members struct {
     name string
}

func printmemberlist(memberList [4]members){
    for i := 0; i < len(memberList); i++{
        members:=memberList[i]
        fmt.Println(members.name)
    }
}

func booklist(libraryBooks [4]book) (checkedout string){
    for i := 0; i < len(libraryBooks); i++ {
        book := libraryBooks[i]
        if book.checkedout && book.checkoutby != ""  {
             fmt.Println(book.name, "is checked out by", book.checkoutby)
           checkedout = book.name 
        }else if book.checkedout {
            fmt.Println(book.name, "is checked out")
            checkedout = book.name 
        }else {
            fmt.Println(book.name, "is checked in")
        }
    }
    return checkedout
}
func main() {
    memberList := [...]members{
        {name: "Julie"},
        {name: "Sally"},
        {name:"Fred"},
        {name:"Kyle"},}
    printmemberlist(memberList)
    libraryBooks := [...]book{
		       {name: "Harry Potter"},
		       {name: "Eragon"},
		       {name: "Sun Stand Still"},
		       {name: "Terraform In Action"},
            }
    booklist(libraryBooks)
    libraryBooks[2].checkedout = true
	libraryBooks[3].checkedout = true
    libraryBooks[3].checkoutby = "Julie"
    booklist(libraryBooks)
}
