package checker

func RunChecks(cfg Config) []CheckResult {
    total := len(cfg.Websites) + len(cfg.APIs)
    results := make([]CheckResult, 0, total)
    ch := make(chan CheckResult, total)

    // Run websites concurrently
    for _, w := range cfg.Websites {
        go func(w Website) {
            ch <- CheckWebsite(w.Name, w.URL)
        }(w)
    }

    // Run APIs concurrently
    for _, a := range cfg.APIs {
        go func(a API) {
            ch <- CheckAPI(a.Name, a.URL, a.Content)
        }(a)
    }

    // Collect results
    for i := 0; i < total; i++ {
        results = append(results, <-ch)
    }

    return results
}