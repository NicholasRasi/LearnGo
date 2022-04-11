# Packages and Imports
> A package is a focused area of functionality

In a Go program, each directory is associated with a unique package.

- importing a package gives access to everything in the package
- ``import "fmt"`` or
```go
import (
    "fmt",
    str "strings"
)
```
- the functions inside the package can be referenced with the name of the package
and the name of the item, e.g. ``fmt.Println``
- the module name can be aliased

### Create a Package
- a package is a folder
- functions with the first character uppercase are public
- If you want to make something package visible, only make the first character in the name of the item lowercase
- relative imports are bad practice


### Third Party Packages
- ``dep`` is used to manage the vendor directory
- ``dep init`` and ``dep ensure`` are useful commands for dependency management
- Vigo is a new project for dependencies management
- a package can be distributed with any public code repository
