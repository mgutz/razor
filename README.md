# razor

`razor` is a Go port of ASP.NET's Razor view engine.  `razor` is a code
generator which compiles Razor templates into Go functions. It is fast, type
safe and escapes all values by default.

# Usage

Install

```sh
go get github.com/mgutz/gorazor
```

Running

```sh
gorazor template_folder output_folder
gorazor template_file output_file
```

## Layout & Views

Let's cover the basic use case of a view with a layout. In `razor` each template becomes
a Go function.  A layout is a function a which receives the rendered result of a view.
That is, given a layout function named `Layout` and a view function `View`, the view
is rendered as `Layout(View())`.

Let's step through it. First define a layout, `views/layout/base.go.html`

```html
@{
    +func(title string, css razor.SafeBuffer, body razor.SafeBuffer, js razor.SafeBuffer)
}

<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8" />
	<title>@title</title>
        @css
</head>
<body>
        <div class="container">@body</div>
        @js
</body>
</html>
```

The first block of template instructs `razor` how to generate the function. In
this example, the header declares a function with a signature of

    (title string, css razor.SafeBuffer, body razor.SafeBuffer, js razor.SafeBuffer)

Notice the arguments are used in the template as variables denoted by `@`.
The layout expects these values from the view.

Let's now define a view `views/index.go.html` to use the layout.

```html
@{
    import (
        "views/layout"
    )

    // `+` indicates a directive and is intentionally not valid Go syntax
    +func (name string)
    +return layout.Base(title, "", VIEW, js())

    // inline code inserted in the function
    title := "Welcome Page"
}

<h2>Welcome to homepage</h2>

@section js {
<script>
    alert('hello! @name')
</script>
}
```

This view has a signature of `(name string)` which means a `name` value must be passed in
as an argument.  A variable `title` is set in a code block and is used by the layout.
A `section` named `js` becomes its own function. The magic all happens in the
function's return value of `layout.Base(title, "", VIEW, js())`. `VIEW` is a placeholder
for the rendered value of the view template.

To call from Go code

```go
import (
    "views"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
    views.Index("gopherito").WriteTo(w)
}
```

See [working example](example).

| Description | Template | Generated code |
| ------------| -------- | ---------------|
| View |  [index.go.html](example/views/index.go.html) | [index.go](example/views/index.go) |
| Layout | [default.go.html](example/views/layout/default.go.html) | [default.go](example/views/layout/default.go) |

## Syntax

### Variable

* `@variable` to insert **string** variable into html template
* variable could be wrapped by arbitrary go functions
* variable inserted will be automatically [escaped](http://golang.org/pkg/html/template/#HTMLEscapeString)

```html
<div>Hello @user.Name</div>
```

```html
<div>Hello @strings.ToUpper(req.CurrentUser.Name)</div>
```

`razor` escapes any value that is not of type `razor.SafeBuffer`. To
insert unescaped data use `github.com/mgutz/gorazor/html#Raw`

```html
    @html.Raw("<h2>Heading 2</div>")
```

### Helper Functions

To create function whose result should not be escaped, return
`razor.SafeBuffer`. Here's how `Raw` is implemented.

```go
func Raw(t interface{}) razor.SafeBuffer {
    buffer := razor.NewSafeBuffer()
    buffer.WriteString(fmt.Sprint(t))
    return buffer
}
```

### Flow Control

```php
@if condition {
    ...
}

@if condition {
    ...
} else {
    ...
}

@for condition {
    ...
}

@{
    switch variable {
    case 1:
          <p>...</p>
    case 2:
          <p>...</p>
    default:
          <p>...</p>
    }
}
```

### Code block

It's possible to insert arbitrary go code block in the template, like create new variable.

```html
@{
    username := u.Name
    if u.Email != "" {
        username += "(" + u.Email + ")"
    }
}
<div class="welcome">
<h4>Hello @username</h4>
</div>
```

### Declaration

The **first code block** is strictly for declaration:

```
@{
    import ...
    +func ...
    +return ...
}
```

*   **import** - Optional. Import packages used by view.

        import (
            "sipin/views"
            "sipin/models"
        )

*   **+func** - Optional. Declare the signature for the generated function. Defaults to `()`

        +func (user *models.user)

*   **+return** - Optional. Override the return value. Defaults to rendered template value.

        +return views.Layout(VIEW, scripts())


**first code block** must be at the beginning of the template, i.e. before any html.

`import` must be wrapped in `()`


### Helper / Include other template

`razor` compiles templates to go functions. Composition and helpers are simply
Go functions which return values that can be converted to `string`.

If your helper needs to write unescaped values to the output buffer, use
`razor.SafeBuffer` which is a `bytes.Buffer`. `@razor.Raw` may also be used but
is not recommended. Keep your template clean by returning `SafeBuffer`.

## Conventions

*   Package name is derived from directory name.

        "views/layout" => package layout
        "views/home" => package home

*   Template filename must have the extension name `.go.html`

*   Function name is the Capitalized basename of the file without the extension.

        "views/layout/default.html" => function Default()
        "views/home/index.gothml" => function Index()

## FAQ

## Watch go.html files?

Use [gosu](http://github.com/mgutz/gosu).  See `example` directory on how
easy it is to use *gosu*

# Credits

The original and likely more awesome [sipin gorazor](https://github.com/sipin/gorazor).


