package reporter

import (
    "api-monitor/internal/checker"
    "io/ioutil"
    "regexp"
    "fmt"
)

// Update README between <!-- STATUS_START --> and <!-- STATUS_END -->
func UpdateReadme(results []checker.CheckResult, path string) error {
    contentBytes, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }
    content := string(contentBytes)

    table := "| Name | Status | Response (ms) |\n|------|--------|---------------|\n"
    for _, r := range results {
        emoji := "🟢"
        if r.Status == "WARN" { emoji = "🟡" }
        if r.Status == "DOWN" { emoji = "🔴" }
        table += fmt.Sprintf("| %s | [%s](%s) | %s %s | %dms |\n", r.Name, "Link", r.URL, emoji, r.Status, r.ResponseMS)
    }

    re := regexp.MustCompile(`(?s)<!-- STATUS_START -->.*?<!-- STATUS_END -->`)
    newContent := ""
    if re.MatchString(content) {
        newContent = re.ReplaceAllString(content, fmt.Sprintf("<!-- STATUS_START -->\n%s<!-- STATUS_END -->", table))
    } else {
        newContent = content + "\n<!-- STATUS_START -->\n" + table + "<!-- STATUS_END -->"
    }

    return ioutil.WriteFile(path, []byte(newContent), 0644)
}