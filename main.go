package main

import (
  "fmt"
  "sync"
  "time"
)

/*
Three kinds of threads share access to a singly-linked list: searchers, inserters and deleters. Searchers merely examine the list; hence they can execute concurrently with each other.

Inserters add new items to the end of the list; insertions must be mutually exclusive to preclude two inserters from inserting new items at about the same time.

However, one insert can proceed in parallel with any number of searches. Finally, deleters remove items from anywhere in the list. At most one deleter process can access the list at a time, and deletion must also be mutually exclusive with searches and insertions.
*/

func searcher(search_wg *sync.WaitGroup, delete_wg *sync.WaitGroup) {
  delete_wg.Wait()
  search_wg.Add(1)
  
  fmt.Println("searching....")
  time.Sleep(5*time.Second)

  search_wg.Done()

  fmt.Println("search done...")
  time.Sleep(5*time.Second)
}

func inserter(im *sync.Mutex, insert_wg *sync.WaitGroup, delete_wg *sync.WaitGroup) {
  delete_wg.Wait()
  insert_wg.Add(1)

  im.Lock()

  fmt.Println("inserting.... ")

  time.Sleep(5*time.Second)

  im.Unlock()
  insert_wg.Done()

  fmt.Println("inserting done...")
  time.Sleep(5*time.Second)
}

func deleter(dm *sync.Mutex, search_wg *sync.WaitGroup, delete_wg *sync.WaitGroup, insert_wg *sync.WaitGroup) {
  search_wg.Wait()
  insert_wg.Wait()

  //delete_wg.Wait()

  delete_wg.Add(1)
  dm.Lock()

  fmt.Println("Deleting....")

  time.Sleep(5*time.Second)

  fmt.Println("asdfasdfasd")

  dm.Unlock()
  
  delete_wg.Done()
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
  
  var n int = 5 // 5 of each type of thread
  
  search_wg := &sync.WaitGroup{}
  insert_wg := &sync.WaitGroup{}
  delete_wg := &sync.WaitGroup{}
  im :=        &sync.Mutex{}
  dm :=        &sync.Mutex{}

    go searcher(search_wg, delete_wg)
    go searcher(search_wg, delete_wg)
    go inserter(im, insert_wg, delete_wg)
    go inserter(im, insert_wg, delete_wg)
    go deleter(dm, search_wg, delete_wg, insert_wg)
    go deleter(dm, search_wg, delete_wg, insert_wg)

  fmt.Println(n)

  wg := &sync.WaitGroup{}
  wg.Add(1)
  wg.Wait()
}
