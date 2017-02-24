package main

import (
    "fmt"
    "math/rand"
    "net/http"
    // "io/ioutil"
    // "reflect"
    // "sync"
    "github.com/gin-gonic/gin"
    // "time"
    "github.com/NaySoftware/go-fcm"
    "strconv"
    "log"
    "github.com/robfig/cron"
)

const (
     serverKey = "YOUR-KEY"
     rootTopic = "/topics/fooTopic"
     MIN_THRESHOLD=0
     MAX_THRESHOLD=16
)

var client = &http.Client{}


type Subscription struct {
    Topic   string `form:"topic" json:"topic" binding:"required"`
    InstanceId  string `form:"instance_id" json:"instance_id" binding:"required"`
}

// construct funcs for each interval
// publishes to root_topic//stream/hour
func mkPub(utc_hour int, root_topic string) func() {
  topic := root_topic + "-" + strconv.Itoa(utc_hour)
  return func() {

    c := fcm.NewFcmClient(serverKey)
    r := rand.New(rand.NewSource(99))
    random_num := r.Int31()
    data := map[string]string{
        "random_num": strconv.Itoa(int(random_num)),
    }
    c.NewFcmMsgTo(topic, data)

    status, err := c.Send()

    if err == nil {
        status.PrintResults()
    } else {
        fmt.Println(err)
    }
  }
}

func SubscribeToTopic(c *gin.Context) {
    var json Subscription
    if c.BindJSON(&json) == nil {
        // build the request object
        req, err := http.NewRequest("GET",
            "https://iid.googleapis.com/iid/v1/" +
            json.InstanceId+"/rel/topics/"+json.Topic, nil)
        req.Header.Add("Authorization", "key="+serverKey)
        req.Header.Add("Content-Type", "application/json")
        req.Header.Add("Content-Length", "0")
        resp, err := client.Do(req)
        if err != nil {
          fmt.Println(err)
        } else {
            fmt.Println(resp)
        }
        c.String(200, "hello")
        log.Println("hallo")
        defer resp.Body.Close()
    }
}



func main() {
  r := gin.New()
  r.Use(gin.Logger())
  r.Use(gin.Recovery())
  c := cron.New()
  for i := 0; i < 24; i++ {
    // test := mkPub(i, rootTopic)
    // c.AddFunc("" + strconv.Itoa(i) + " * * * * *", test)
  }
  c.AddFunc("@every 10s",      func() { fmt.Println("Every 10s") })
  v1 := r.Group("api/v1")
  {
      v1.POST("/subscribe", SubscribeToTopic)
  }
  c.Start()
  r.Run(":8080")
}
