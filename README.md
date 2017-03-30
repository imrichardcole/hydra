Hydra
=====

A simple go utility to run command line processes in parallel.

Usage
=====

Create a job file (example.job):

  {
    "output" : true,
    "commands" : {
      "ls" : "-lrt",
      "ps" : "-al"
    }
  }

Start Hydra refering to this file:

  ./hydra example.job

To turn off output to standard out change the "output" setting to false:

{
  "output" : false,
  "commands" : {
    "ls" : "-lrt",
    "ps" : "-al"
  }
}
