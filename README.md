# razor

**razor** is a Go port of ASP.NET's Razor view engine with less magic.
**razor** is a code generator which compiles Razor templates into a Go package of template functions.
**razor** is fast, reflection-less and escapes all values by default.

## Usage

Install

```sh
go get -u github.com/mgutz/razor/cmd/razor
```

Running

```sh
razor <folder or file> <output folder or file>
```

## Layout & Views

Let's cover the basic use case of a view with a layout. In **razor** each template becomes
a Go function. A layout is a function a which receives the rendered result of a view.
That is, given a layout function named `Layout` and a view function `View`, the view
is rendered as `Layout(View())`.

Let's step through it. First define a layout, `views/layout/base.go.html`

```html
@{
    +func(title string, body *razor.SafeBuffer, sections razor.Sections)
}

<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>@title</title>
    <link rel="stylesheet" href="/@locals["version"]/css/style.css">
    @sections["css"]
</head>
<body>
    <div class="container">@body</div>
    @sections["js"]
</body>
</html>
```

The first block of the template instructs **razor** how to generate the function.
The generated function looks like this

```go
// from dir
package "layout"

// from +func
func Base(title string, body *razor.SafeBuffer, sections razor.Sections) *razor.SafeBuffer {
    _buffer := razor.NewSafeBuffer()
    locals := razor.Locals

    // ... markup written to _buffer

    return _buffer
}
```

Key points

1.  The package name is derived from the directory.
2.  The generated function is derived from `+func` directive
3.  **razor** adds a `locals` variable accessible as `@locals`.
    Call `razor.SetLocals` once to initialize locals for all templates.
4.  A section in `sections` and `body` are of type `*razor.SafeBuffer`

Let's now define a view `views/index.go.html` function to use the layout.

```html
@{
    +import (
        "views/layout"
    )
    +func (name string)
    title := "Welcome"
    +return layout.Base(title, VIEW, SECTIONS)
}

<h2>Welcome to homepage</h2>

@section js {
    <script>
        alert('hello! @name')
    </script>
}
```

Key points

The `+return` directive instructs razor to call `layout.Base` with arguments `title`, `VIEW`
and `SECTIONS` which corresponds to the parameters defined in the layout.

There are two reserved keywords which may be used in the return statement

- `VIEW`:  the rendered buffer of the view, ie everything outside of sections aka body
- `SECTIONS`: a map of rendered buffer sections by name


## Using Generated Package

To call from Go code

```go
import (
    "views"
    "shared"
    "github.com/mgutz/razor/razor"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "Foo"}
	views.Index(user).WriteTo(w)
}

func main() {
	razor.SetLocals(razor.M{
		"version": "1.0.0",
	})

	http.HandleFunc("/", viewHandler)
	http.Handle("/{{version}}/", http.FileServer(http.Dir("public")))
	port := ":8080"
	fmt.Printf("Browse 127.0.0.1%s\n", port)
	http.ListenAndServe(":8080", nil)
}
```

## Example

See [working example](example).

| Description | Template | Generated code |
| ------------| -------- | ---------------|
| View |  [index.go.html](example/views/index.go.html) | [index.go](example/views/index.go) |
| Layout | [default.go.html](example/views/layout/default.go.html) | [default.go](example/views/layout/default.go) |


To build

    gosu views

To watch

    gosu views --watch
