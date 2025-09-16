package checker

import (
    "bytes"
    "encoding/json"
    "net/http"
    "time"
)

func CheckAPI(name, url string, content map[string]interface{}) CheckResult {
    start := time.Now()
    status := "UP"

    var bodyBytes []byte
    var err error
    if content != nil {
        bodyBytes, err = json.Marshal(content)
        if err != nil {
            status = "DOWN"
        }
    }

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
    req.Header.Set("Content-Type", "application/json")

    client := http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    duration := int(time.Since(start).Milliseconds())

    if err != nil || resp.StatusCode >= 400 {
        status = "DOWN"
    }

    return CheckResult{
        Name:       name,
        URL:        url,
        Status:     status,
        ResponseMS: duration,
    }
}