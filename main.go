package main

import (
    "fmt"
    "math/rand"
    // "io/ioutil"
    // "reflect"
    // "sync"
    "github.com/gin-gonic/gin"
    // "time"
    "github.com/NaySoftware/fcm"
    "strconv"
    "github.com/robfig/cron"
)

const (
     serverKey = "YOUR-KEY"
     rootTopic = "foo-topic"
     MIN_THRESHOLD=0
     MAX_THRESHOLD=16
)

// construct funcs for each interval
// publishes to root_topic//stream/hour
func mkPub(utc_hour int, stream string, root_topic string) func() {
  topic := root_topic + "/" + strconv.Itoa(utc_hour)
  return func() {
    
    c := fcm.NewFcmClient(serverKey)
    r := rand.New(rand.NewSource(99))
    data := r.Int31()
    x := topic + " + " + strconv.Itoa(data.toInt())
    fmt.Println(x)
    c.NewFcmMsgTo(topic, data)

    status, err := c.Send()

    if err == nil {
    status.PrintResults()
    } else {
        fmt.Println(err)
    }
  }
}

func PostUser(c *gin.Context) {
    // The futur code…
}



func main() {
  r := gin.Default()
  c := cron.New()
  for i := 0; i < 24; i++ {
    test := mkPub(i, rootTopic)
    c.AddFunc("0 0 " + strconv.Itoa(i) + " * * *", test)
  }
  c.AddFunc("@every 10s",      func() { fmt.Println("Every 10s") })
  v1 := r.Group("api/v1")
  {
      v1.POST("/users", PostUser)
      v1.GET("/users", PostUser)
      v1.GET("/users/:id", PostUser)
      v1.PUT("/users/:id", PostUser)
      v1.DELETE("/users/:id", PostUser)
  }
  c.Start()
  r.Run(":8080")
  
  // resp, err := http.Get("http://www.arpansa.gov.au/uvindex/realtime/xml/uvvalues.xml")
  // if err != nil {
  //   fmt.Println(err)
  // }
  // defer resp.Body.Close()

  // var q Arpansa_Query;
  // body, err := ioutil.ReadAll(resp.Body)

  // if err != nil {
  //   panic(err.Error())
  // }

  // fmt.Println(reflect.TypeOf(body))
  // fmt.Println(string(body))
	// xml.Unmarshal(body, &q)
  // fmt.Println(q)

  // messages := make(chan int)
  // var wg sync.WaitGroup
  // wg.Add(1)
  // go func() {
  //   defer wg.Done()
  //   publish_worker(0)
  //   messages <- 1
  // }()

  // go func() {
  //     for i := range messages {
  //         fmt.Println(i)
  //     }
  // }()
  // wg.Wait()
}
