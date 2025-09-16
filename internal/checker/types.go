package checker

type Website struct {
    Name string `yaml:"name"`
    URL  string `yaml:"url"`
}

type API struct {
    Name    string                 `yaml:"name"`
    URL     string                 `yaml:"url"`
    Content map[string]interface{} `yaml:"content"`
}

type CheckResult struct {
    Name       string
    URL        string
    Status     string
    ResponseMS int
}

type Config struct {
    Websites []Website `yaml:"websites"`
    APIs     []API     `yaml:"apis"`
}
