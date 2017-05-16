package main

import (
  "fmt"
)

/*
Three kinds of threads share access to a singly-linked list: searchers, inserters and deleters. Searchers merely examine the list; hence they can execute concurrently with each other.

Inserters add new items to the end of the list; insertions must be mutually exclusive to preclude two inserters from inserting new items at about the same time.

However, one insert can proceed in parallel with any number of searches. Finally, deleters remove items from anywhere in the list. At most one deleter process can access the list at a time, and deletion must also be mutually exclusive with searches and insertions.
*/

func searcher() {

}

func deleter() {

}

func inserter() {

}

func main() {
  list := new(LinkedList)
  list.PushEnd(0)
  list.PushEnd(1)
  list.PushEnd(2)
  list.Delete(0)
  list.Delete(1)
  list.Print()
  fmt.Println("done")
}
