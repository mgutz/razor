
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
`*razor.SafeBuffer`. Here's how `Raw` might be implemented.

```go
func Raw(t interface{}) *razor.SafeBuffer {
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

*   **+import** - Optional. Import additional packages used by  theview.

        +import (
            "sipin/views"
            "sipin/models"
        )

*   **+func** - Optional. Declare the signature for the generated function. Defaults to `()`

        +func (user *models.user)

*   **+return** - Optional. Override the return value. Defaults to rendered template value.

        +return views.Layout(VIEW, scripts())


**first code block** must be at the beginning of the template, i.e. before any html.


### Helper / Include other template

`razor` compiles templates to go functions. Composition and helpers are simply
Go functions which return values that can be converted to `string`.

If your helper needs to write unescaped values to the output buffer, use
`*razor.SafeBuffer` which is a `bytes.Buffer`. `html.Raw` may also be used but
is not recommended. Keep your template clean by returning `*razor.SafeBuffer`.

## Conventions

*   Package name is derived from directory name.

        "views/layout" => package layout
        "views/home" => package home

*   Template filename must have the extension name `.go.html`

*   Function name is the Capitalized basename of the file without the extension.

        "views/layout/default.html" => function Default()
        "views/home/index.gothml" => function Index()

## FAQ

### Watch go.html files?

Use [gosu](http://github.com/mgutz/gosu).  See `tasks` directory on how
easy it is to use *gosu* to build and watch views.

### How to set locals?

Locals is map that is available to all templates. Store simple common data such
as version into the map.  Use `razor.SetLocals()` before rendering any views.
See example project.

# Credits

The original and likely more awesome [sipin gorazor](https://github.com/sipin/gorazor).


