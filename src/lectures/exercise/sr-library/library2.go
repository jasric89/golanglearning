package main

import "fmt"

func printmemberlist(memberList []string) {
    members := make(map[string]int)
    for _, ml:= range memberList {
        members[ml]++
        fmt.Println(ml)
    }
}

func main() {
    memberList := []string{"julie","sally","fred","kyle"}
    printmemberlist(memberList)
}