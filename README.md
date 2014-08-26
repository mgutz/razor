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
        views.Render("index", "Joe").WriteTo(w)
        // or views.Index("Joe").WriteTo(w)
    }
    http.ListenAndServe(":8080", nil)
}
```


There are two ways to render a template

1.  By name. This is useful when the name of the view is dynamic such as a value from a URL segment.
    This function is defined in `razor_render.go` in each package.

        `pkg.Render("index", data).WriteTo(w)`

2.  By func (slightly faster). `pkg.Index(data).WriteTo(w)`


The default code generation accepts a single `interface{}` argument and performs a type assertion like this.

    // +params (user *model.User)
    func GeneratedFunc(__data interface) *razor.SafeBuffer {
        user := __data.(*model.User)
        ...
    }

This is necessary for the `Render(name string, data interface{})` function to work. If you run
`razor --strong SOME_DIR` no type assertions are used and unlimited arguments are allowed.
It will be faster too but in reality `Render` is more practical.

    // +params (user *model.User, other string)
    func GeneratedFunc(user *model.User, other string) *razor.SafeBuffer {
        ...
    }


## Why

Why use Razor over the standard `"html/template"`? It depends.

**razor**

-   Speed, between 3-3.5x faster (see `benchfiles`)
-   Templates become functions
-   Use any package of helper functions directly in templates.
    Note that all inserted values are escaped unless it returns a *razor.SafeBuffer.
-   Use `go` syntax for everything.
    `@for hobby := range hobbies {` instead of `{{ range $hobby := .Hobbies }}`
-   Less reflection. Reflection is used only for HTML escaping.
-   Compilation performed outside of code (watch and server reload with gosu)

**html/template**

-   Fast enough
-   Standard

## Benchmarks

See `benchfiles/` directory. Run with `go test -bench=.`

On 2011 i5 Macbook Pro

    BenchmarkGoTemplate     100000     31107 ns/op
    BenchmarkRazorByName    200000      9092 ns/op
    BenchmarkRazorByFunc    200000      8243 ns/op

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
