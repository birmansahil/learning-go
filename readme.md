**Packages**
Every Go program is made up of packages.

Programs start running in package main.

To use the packages you need to import them into your program. For example:
import “fmt”

**Imports**
To import multiple packages you can either write multiple import statements like:
import “fmt”
import “math”

You can also group the imports into a parenthesized/factored import statement like:
import (
“fmt”
“math”
)

Note: It is good to use the parenthesized import statement.

**Exported Names**
In GO, a name is exported if it begins with a capital letter. For instance “Pi” is exported from the math package.

“pi” does not start with a capital letter, so it is not exported.

You can only refer to exported names of a package after importing it. Any “unexported” names will result in error “undefined: package.exported-name”

fmt.Println(math.pi) - will result in error
fmt.Println(math.Pi) - will print the value of pi

**Functions**
A function can take zero or more arguments.
The data type of arguments comes after their names

func multiply(a int, b int) int {
return a\*b
}

Note: When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

func multiply(a int, b int) int {...}
Can be re-written as
func multiply(a, b int) int {...}

Note: A function can return any number of results and they can be of different data types.

A return value can be named. If there are named return values, they will be treated as variables defined at the top of the function.

Named returns are used to document the meaning of return values.

**Naked Return**
A return statement without arguments returns the named return values. This is known as “naked” return.

Naked return statements should be used only in small functions as they can harm readability in longer functions.

**Variables**
Variables can be declared using the “var” keyword. Variables can be declared at package or function level.

var name datatype

A var declaration can include initializers, one per variable.
var name datatype = value

Note: If an initializer is present, the data type can be omitted. The variable will automatically take the type of initializer.

**Shorthand Variable Declaration**
Inside a function, a shorthand can be used to declare and initialise a variable.

name := value

Note: this shorthand is not available outside a function.
