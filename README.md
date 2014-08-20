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
The generation function looks like this

```go
// from dir
package "layout"

// from +func
func Base(title string, body * razor.SafeBuffer, sections razor.Sections) *razor.SafeBuffer {
    _buffer := razor.NewSafeBuffer()
    locals := razor.Locals

    // ... markup written to _buffer

    return _buffer
}
```

Notice the arguments are used in the template as variables denoted by `@`.
`@locals` is a special variable added by **razor** which is a map
that can initialized by you.

Let's now define a view `views/index.go.html` function to use the layout.

```html
@{
    +import (
        "views/layout"
    )
    +func (name string)
    title := "Welcome Page"
    +return layout.Base(title, VIEW, SECTIONS)
}

<h2>Welcome to homepage</h2>

@section js {
<script>
    alert('hello! @name')
</script>
}
```

This view has a function signature of `(name string)` which means a `name` value must be passed in
as an argument. A variable `title` is defined and becomes an argument for the layout.
The return value matches the signature as expected by layout.

There are two reserved keywords for use in the return statement

- `VIEW`:  the rendered buffer of the view
- `SECTIONS`: a map of rendered sections by name

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
