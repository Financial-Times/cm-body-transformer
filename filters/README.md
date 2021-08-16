# Content Body Filters

`filters` package is a collection of helper functions for text manipulation.

Multiple services in Content & Metadata pipeline require "clean" content body.
Clear of specific data like quotes and tables and all HTML tags.

The aim of this package is to aggregate this cleanup functionality in a single place, and provide easy to reuse interface.

## Install

```shell
go get github.com/Financial-Times/cm-body-transformer
```
## Usage

For more examples see `example_test.go`
```go

import "github.com/Financial-Times/cm-body-transformer/filters"

...

func foo() string {
    body := `<p>content body</p>`
    result := filters.Apply(body, filters.DefaultContentFilters()...)
    return result
}
```
