# API & Website Monitoring Tool

This project is a **Go-based monitoring tool** that checks the status and performance of websites and APIs on a scheduled basis using **GitHub Actions**. It is designed to run **headless browser checks** for websites, detect **console errors**, test APIs with optional request bodies, and report results **concurrently**. Results are automatically updated in the **README** and the **GitHub Actions workflow summary**.

---

<!-- STATUS_START -->
| Name | Status | Response (ms) |
|------|--------|---------------|
| Personal Website | 🟢 UP | 1119 |
| Commute Rentals | 🟢 UP | 1230 |
| Rental Listing API | 🟢 UP | 4686 |
<!-- STATUS_END -->

## Features

- **Website Monitoring**
  - Uses headless Chrome via [Chromedp](https://github.com/chromedp/chromedp) to render pages fully.
  - Detects console errors (JavaScript issues) and page availability.
  - Measures response time.

- **API Monitoring**
  - Supports GET and POST requests with configurable JSON payloads.
  - Validates response status and performance.
  
- **Concurrent Execution**
  - Runs all website and API checks concurrently for faster results.

- **GitHub Integration**
  - Updates a status table in the README while preserving other content.
  - Updates workflow summary in GitHub Actions, triggering native notifications if you watch the repository.

---

## How It Works

1. **Configuration**
   - Define websites and APIs to monitor in `configs/config.yaml`.
   - Example:

```yaml
websites:
  - name: MySite
    url: https://example.com

apis:
  - name: My API
    url: https://exampleapi.com
    content:
        param: content
```

2. **Running the Monitor**

   * The entry point is `cmd/monitor/main.go`.
   * It loads the configuration, runs all checks concurrently, and sends results to the reporter.

3. **Website Checks**

   * Websites are loaded headlessly with Chromedp.
   * Console errors are detected, and page HTML is fetched.
   * Status is marked as:

     * `UP` → page loads with no console errors.
     * `WARN` → page loads but has console errors.
     * `DOWN` → page failed to load.

4. **API Checks**

   * Supports GET or POST requests depending on the presence of a JSON payload in `config.yaml`.
   * Records response time and success/failure.

5. **Reporting**

   * **README**:

     * Updates only the **status table** between the markers:
* **Workflow Summary**:

  * Updates the GitHub Actions summary for each run, visible in the workflow UI.

---

## Usage

1. Install dependencies:

```bash
go mod tidy
```

2. Run locally:

```bash
go run ./cmd/monitor
```

3. Schedule weekly checks with **GitHub Actions**:

* See `.github/workflows/monitor.yml`.
* Automatically runs, updates the README, and triggers GitHub notifications.

---

## Dependencies

* [Go](https://golang.org/)
* [Chromedp](https://github.com/chromedp/chromedp)
* [yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3)

---

## Notes

* Only the status table in the README is overwritten; all other content is preserved.
* For headless browser checks, GitHub Actions runs Chrome in headless mode automatically.
* The workflow can be manually triggered via the Actions tab or scheduled with a cron expression.

---