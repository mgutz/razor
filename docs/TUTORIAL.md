# Tutorial

## Layout & Views

Let's cover the basic use case of a view with a layout. In **razor** each template becomes
a Go function. A layout is a function a which receives the rendered result of a view.
That is, given a layout function named `Layout` and a view function `View`, the view
is rendered as `Layout(View())`.

Let's step through it. First define a layout, `views/layout/base.go.html`

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
    @RenderSection("css")
</head>
<body>
    <div class="container">@RenderBody()</div>
    @RenderSection("js")
</body>
</html>
```

The first block of the template instructs **razor** how to generate the function.
The generated function looks like this

```go
package "layout"

func Base(title string, body *razor.SafeBuffer, sections razor.Sections) *razor.SafeBuffer {
    _buffer := razor.NewSafeBuffer()
    App := razor.App
    RenderBody := func() *razor.SafeBuffer {
        return body
    }
    RenderSection := func(section string) *razor.SafeBuffer {
        return sections[section]
    }

    // ... markup written to _buffer

    return _buffer
}
```

Key points

1.  The package name is derived from the directory.
2.  The function name is the basename of the template.
3.  The generated function signature is derived from `+params` directive
    and always has a return value of `*razor.SafeBuffer`
4.  params `...` means insert body and section variables
4.  **razor** adds an `App` variable accessible as `@App` representing app-wide state.
    Call `razor.SetAppState` once to initialize the `App` map.

Let's now define a view `views/index.go.html` function to use the layout.

```html
@{
    +import (
        "views/layout"
    )
    +params (name string)
    +return layout.Base("Welcome " + name, ...)
}

<h2>Welcome to homepage</h2>

@section js {
    <script>
        alert('hello! @name')
    </script>
}
```

Key points

The `+return` directive instructs razor to call `layout.Base` with a title argument and
the `...` means insert rendered view and section variables here.

## Using Generated Package

To call from Go code

```go
import (
    "views"
    "models"
    "github.com/mgutz/razor"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "Foo"}
	views.Index(user).WriteTo(w)
}

func main() {
	razor.SetAppState(map[string]interface{}{
		"version": "1.0.0",
	})

	http.HandleFunc("/", viewHandler)
	http.Handle("/{{version}}/", http.FileServer(http.Dir("public")))
	port := ":8080"
	fmt.Printf("Browse 127.0.0.1%s\n", port)
	http.ListenAndServe(":8080", nil)
}
```
