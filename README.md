# vanta-sdk-go

[![Go Report Card](https://goreportcard.com/badge/github.com/adrianosela/vanta-sdk-go)](https://goreportcard.com/report/github.com/adrianosela/vanta-sdk-go)
[![Documentation](https://godoc.org/github.com/adrianosela/vanta-sdk-go?status.svg)](https://godoc.org/github.com/adrianosela/vanta-sdk-go)
[![GitHub issues](https://img.shields.io/github/issues/adrianosela/vanta-sdk-go.svg)](https://github.com/adrianosela/vanta-sdk-go/issues)
[![license](https://img.shields.io/github/license/adrianosela/vanta-sdk-go.svg)](https://github.com/adrianosela/vanta-sdk-go/blob/master/LICENSE)

Unofficial Vanta API Client for Go.

Implementing only what I need for my current use case. Feel free to open issues for endpoints you need or open a PR.

## Example

```
ctx := context.Background()

v, err := vanta.New(
    ctx,
    vanta.WithOAuthCredentials(
        os.Getenv("VANTA_OAUTH_CLIENT_ID"),
        os.Getenv("VANTA_OAUTH_CLIENT_SECRET"),
    ),
)
// check err

output, err := v.ListVulnerabilities(
    ctx,
    vanta.WithPageSize(100),
    vanta.WithSearchQuery("openssl"),
    vanta.WithIsFixAvailable(true),
    vanta.WithSeverity(vanta.VulnerabilitySeverityCritical),
    vanta.WithSLADeadlineBeforeDate(time.Now().Add(time.Hour*24*30)),
)
// check err

for _, vulnerability := range output.Results.Data {
    // do something useful
}
```