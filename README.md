# cron-registry
Easily register all of your cron jobs in one place.

## Creating A Registry
```go
package main

import (
    "fmt"

    . "github.com/luckfire-go/cron-scheduler"
)

func main() {
    registry := NewRegistry()

    registry.AddJobs([]RegistryItem{
        {
            Enabled: true,
            Interval: "@every 1m",
            TaskFunc: func() {
                fmt.Println("Hello, world!")
            },
        },
    })

    registry.Start()
}
```

## Add Success / Add Failure Events
```go
package main

import (
    "fmt"

    . "github.com/luckfire-go/cron-scheduler"
)

func main() {
    registry := NewRegistry()

    registry.OnJobAddSuccess = func(job *RegistryItem) {
        fmt.Println("Job added successfully:", job.Name())
    }
    registry.OnJobAddFailure = func(job *RegistryItem, err error) {
        fmt.Println("Job failed to add:", job.Name(), err)
    }

    registry.AddJobs([]RegistryItem{
        {
            Enabled: true,
            Interval: "@every 1m",
            TaskFunc: func() {
                fmt.Println("Hello, world!")
            },
        },
    })

    registry.Start()
}
```