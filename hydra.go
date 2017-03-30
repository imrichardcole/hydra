package main

import (
  "fmt"
  "bytes"
  "os/exec"
  "os"
  "encoding/json"
  "time"
)

type Job struct {
    Output bool
    Commands map[string] string
}

func LaunchCommand(command, args string, output bool) int {
  cmd := exec.Command(command, args)
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    panic(err)
  }
  if output {
    fmt.Println("result of", command, "was", out.String())
  }
  return 1
}

func LaunchJob(job Job) int {
  for command, args := range job.Commands{
    go LaunchCommand(command, args, job.Output)
  }
  fmt.Println(len(job.Commands), "commands launched, sleeping to allow time to finish")
  time.Sleep(time.Millisecond * 3000)
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
  args := os.Args[1:]
  jobfile := args[0]
  fmt.Println("running jobs from", jobfile)
  job := LoadJobFile(jobfile)
  LaunchJob(job)
}
