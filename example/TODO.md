@{
  +import (
    "github.com/mgutz/razor/example/models"
    "github.com/mgutz/razor/example/views"
  )
  +layout Layout

  ViewData["title"] = "title": "Razor + Go = love"
}

@{
  +is_layout
  +import (
    "github.com/mgutz/razor/example/models"
    "github.com/mgutz/razor/example/views"
  )
}

views.Render("admin/index", razor.ViewData{})
