package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "strconv"
)

const DEFAULT_PORT = 8080

func getPort() int64 {
  port, _ := strconv.ParseInt(os.Getenv("PORT"), 10, 32)
  if port == 0 {
    fmt.Println("PORT env var not set, running on default port:", DEFAULT_PORT)
    port = DEFAULT_PORT
  } else {
    fmt.Println("running on port:", port)
  }

  return port
}

// https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
func redirect(w http.ResponseWriter, req *http.Request) {
  // remove/add not default ports from req.Host
  target := "https://" + req.Host + req.URL.Path
  if len(req.URL.RawQuery) > 0 {
    target += "?" + req.URL.RawQuery
  }
  log.Printf("redirect to: %s", target)
  http.Redirect(w, req, target,
    // see @andreiavrammsd comment: often 307 > 301
    http.StatusTemporaryRedirect)
}

func main() {
  http.ListenAndServe(fmt.Sprintf(":%d", getPort()), http.HandlerFunc(redirect))
}
