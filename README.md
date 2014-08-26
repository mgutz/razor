# razor

**razor** is a CLI code generator to compile [Razor-like](http://www.asp.net/web-pages/tutorials/basics/2-introduction-to-asp-net-web-programming-using-the-razor-syntax)
templates into go functions.  **razor** is fast and escapes all values by default.

On 2012 i5 Macbook Pro. See [benchfiles](benchfiles) directory

    BenchmarkGoTemplate   100000   28250 ns/op    3712 B/op    60 allocs/op
    BenchmarkRazorByName  500000    7409 ns/op    2533 B/op    18 allocs/op
    BenchmarkRazorByFunc  500000    7332 ns/op    2533 B/op    18 allocs/op

Layout `views/layout.go.html`

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

Run from terminal

```sh
razor .
```

Start server

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


## Why

**razor**

-   Faster, roughly 3-3.5x faster (see `benchfiles`)
-   Compiles as go functions. Spot errors quickly.
-   Use any package of helper functions directly in templates.
    Note that all inserted values are escaped unless it returns a *razor.SafeBuffer.
-   Familiar go syntax

        <!-- razor -->
        @for hobby := range hobbies {
            <p>@hobby</p>
        }

        <!-- html/template $sometimes, .othertimes -->
        {{ range $hobby := .Hobbies }}
            <p>{{ $hobby }}</p>
        {{ end }}

-   Less reflection

**html/template**

-   Fast enough
-   Standard

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
    go get -u github.com/mgutz/gosu/cmd/gosu

    cd $GOPATH/src/github.com/mgutz/razor
    gosu example

Restart server on view change

    gosu example --watch

## Rendering Templates

There are two ways to render a template

1.  By name. This is useful when the name of the view is dynamic such as a value from a URL segment.
    This function is defined in `razor_render.go` in each package.

        views.Render("index", data).WriteTo(w)

2.  By func (slightly faster).

        views.Index(data).WriteTo(w)


The default code generation mode accepts a *single* `interface{}` argument and performs a type assertion
inside the function

    // +params (user *model.User)
    func GeneratedFunc(__data interface) *razor.SafeBuffer {
        user := __data.(*model.User)
        ...
    }

This above signature is necessary for the `Render(name string, data interface{})` function. If you run
`razor --strong SOME_DIR` no type assertions are used and unlimited arguments are allowed.
It is slightly faster too.

    // +params (user *model.User, other string)
    func GeneratedFunc(user *model.User, other string) *razor.SafeBuffer {
        ...
    }


## Example

See [working example](example).

| Description | Template | Generated code |
| ------------| -------- | ---------------|
| View |  [index.go.html](example/views/front/index.go.html) | [index.go](example/views/front/index.go) |
| Layout | [default.go.html](example/views/front/layout.go.html) | [default.go](example/views/front/layout.go) |

## What is implemented?

While **razor** is based on Micorosoft's implementation, **razor** is geared
towards being a template engine with layout support. Microsoft helper functions
are not implemented. Import a helper package directly in a template.

## Credit

This package uses the lexer/parser of [sipin gorazor](https://github.com/sipin/gorazor).
