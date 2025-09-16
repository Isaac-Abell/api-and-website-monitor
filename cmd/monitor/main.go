package main

import (
    "api-monitor/internal/checker"
    "api-monitor/internal/config"
    "api-monitor/internal/reporter"
    "fmt"
)

func main() {
    cfg := config.LoadConfig("configs/config.yaml")
    results := checker.RunChecks(cfg)

    for _, r := range results {
        fmt.Printf("[%s] %s (%s) - %dms\n", r.Status, r.Name, r.URL, r.ResponseMS)
    }

    err := reporter.UpdateReadme(results, "README.md")
    if err != nil {
        fmt.Println("Error updating README:", err)
    }
}
