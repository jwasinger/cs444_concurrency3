package main

import (
  "fmt"
  "time"
  "sync"
  "math/rand"
)

func threadPrint(msg string) {
  t := time.Now()
  t.Format("20060102150405")
  fmt.Printf("%s, %d: %s\n", t.Format("20060102150405"), getGID(), msg)
}

func main() {
  l := &LinkedList{}

  insert_wg := &sync.WaitGroup{}
  delete_wg := &sync.WaitGroup{}
  search_wg := &sync.WaitGroup{}

  im := &sync.Mutex{}
  dm := &sync.Mutex{}

  numSearchers := 2
  numDeleters := 2
  numInserters := 2

  //inserters
  for i := 0; i < numInserters; i++ {
    go func() {
      for ; ; {
        //dont run while deleters are running
        delete_wg.Wait()

        insert_wg.Add(1)

        threadPrint("inserting...")
        time.Sleep(1*time.Second)

        rand_val := rand.Intn(100)
        im.Lock()
        l.PushEnd(rand_val)
        //fmt.Println(l.Size())
        l.Print()
        im.Unlock()

        insert_wg.Done()
        time.Sleep(1*time.Second)
      }
    } ()
  }

  //deleters
  for i := 0; i < numDeleters; i++ {
    go func() {
      for ; ; {
        threadPrint("deleting blocked")

        insert_wg.Wait()
        search_wg.Wait()
        
        delete_wg.Add(1)

        //fmt.Println("deleting....")
        threadPrint("deleting...")
        time.Sleep(time.Second*1)

        dm.Lock()
        l.Delete(l.Size()-1)
        fmt.Println(l.Size())
        dm.Unlock()

        delete_wg.Done()
        time.Sleep(1*time.Second)
      }
    } ()
  }

  //searchers
  for i := 0; i < numSearchers; i++ {
    go func() {
      for ; ; {
        threadPrint("searching blocked")
        delete_wg.Wait()
        search_wg.Add(1)

        threadPrint("searching...")
        time.Sleep(time.Second*1)
        /*
        randIndex := rand.Intn(l.Size()-1)
        found := l.Search(randIndex)
        fmt.Printf("%d: found %d\n", getGID(), found)
        */

        search_wg.Done()
        time.Sleep(1*time.Second)
      }
    } ()
  }

  time.Sleep(5*time.Minute)
}

