package checker

import (
    "context"
    "fmt"
    "time"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/runtime"
)

func CheckWebsite(name, url string) CheckResult {
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    var consoleErrors []string

    chromedp.ListenTarget(ctx, func(ev interface{}) {
        if e, ok := ev.(*runtime.EventConsoleAPICalled); ok {
            for _, arg := range e.Args {
                consoleErrors = append(consoleErrors, fmt.Sprintf("%v", arg.Value))
            }
        }
    })

    start := time.Now()
    err := chromedp.Run(ctx,
        chromedp.Navigate(url),
        chromedp.WaitReady("body", chromedp.ByQuery),
    )
    duration := int(time.Since(start).Milliseconds())

    status := "UP"
    if err != nil {
        status = "DOWN"
    } else if len(consoleErrors) > 0 {
        status = "WARN"
    }

    return CheckResult{
        Name:       name,
        Status:     status,
        ResponseMS: duration,
    }
}