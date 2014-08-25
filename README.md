# razor

**razor** is a code generator which compiles Razor templates into a Go package of template functions.
**razor** is fast and escapes all values by default.
**razor** is a Go port of ASP.NET's Razor view engine with less magic.

Layout (`views/layout.go.html`)

```html
@{
    +params (title string, ...)
}

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>@title</title>
    <link rel="stylesheet" href="/@App["version"]/css/style.css">
</head>
<body>
    <div id="main">@RenderBody()</div>
    @RenderSection("scripts")
</body>
</html>
```

Page `views/index.go.html`

```html
@{
    +params (name string)
    +return Layout("Welcome " + name, ...)
}

<h2>Welcome to homepage</h2>
<p>This is the body</p>

@section scripts {
    <script>
        alert('hello! @name')
    </script>
}
```

To use template

-   Run from terminal

        razor .

```go
import (
    "views"
    "github.com/mgutz/razor"
)

func main() {
    razor.SetAppState(razor.M{
        "version": "1.0.0",
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        views.Index("Joe").WriteTo(w)
    }
    http.ListenAndServe(":8080", nil)
}
```


## Why

Why use Razor over the standard `"html/template"`? It depends.

**razor**

-   Speed, almost 3x faster (see `benchfiles`)
-   Templates become functions
-   Use `go` syntax for everything.
    `@for hobby := range hobbies {` instead of `{{ range $hobby := .Hobbies }}`
-   Less reflection. Reflection is used only for HTML escaping.
-   Compilation performed outside of code (watch and server reload with gosu)

**html/template**

-   Fast enough
-   Standard

## Benchmarks

See `benchfiles/` directory

    BenchmarkGoTemplate     200000     13148 ns/op
    BenchmarkRazor          500000      4636 ns/op

## Usage

Install

```sh
go get -u github.com/mgutz/razor/cmd/razor
```

Running

```sh
razor <folder or file> [output folder or file]
```

Building views efficiently with [gosu](https://github.com/mgutz/gosu)

    # get gosu task runner
    go get -u github.com/mgutz/gosu
    go get -u github.com/mgutz/gosu/cmd/gosu

    cd $GOPATH/src/github.com/mgutz/razor
    gosu example

Restart server on view change

    gosu example --watch

## Example

See [working example](example).

| Description | Template | Generated code |
| ------------| -------- | ---------------|
| View |  [index.go.html](example/views/front/index.go.html) | [index.go](example/views/front/index.go) |
| Layout | [default.go.html](example/views/front/layout.go.html) | [default.go](example/views/front/layout.go) |


## Credit

This package is a fork of [sipin gorazor](https://github.com/sipin/gorazor).
