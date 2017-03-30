package main

import (
  "fmt"
  "bytes"
  "os/exec"
  "os"
  "encoding/json"
)

type Job struct {
    Commands []string
}

func LaunchJob(job Job) int {
  for _, command := range job.Commands {
    cmd := exec.Command(command)
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
      panic(err)
    }
    fmt.Println("result of", command, "was", out.String())
  }
  return 1
}

func LoadJobFile(filename string) Job {
  file, _ := os.Open(filename)
  decoder := json.NewDecoder(file)
  job := Job{}
  err := decoder.Decode(&job)
  if err != nil {
    panic(err)
  }
  return job
}

func main() {
  job := LoadJobFile("example.job")
  LaunchJob(job)
}
