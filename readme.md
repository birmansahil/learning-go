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

**Zero Values**

Variables declared without an explicit initial value are given their zero value.

The zero values are:
0 for numeric types
false for boolean types
“” (empty) for strings

**Basic Data Types**

boolean
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //alias of uint8
rune //alias of int32, a unicode code point
float32 float64
complex64 complex128
The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. Generally you should use int for integer values unless you have a specific reason to use a sized or unsigned integer type.

int - stores positive & negative values
uint - stores non-negative values, providing a wider range of positive numbers compared to int

**Type Conversions**

Go lang enforces strict data types and assignment between different data types requires explicit conversion.

The expression T(v) converts the value v to the type T.

var i int = 10
var f float64 = float64(i)

**Type Inference**

When a variable is declared without specifying an explicit type, the variable’s type is inferred from the value on the right hand side.

a := 14 // int
b := 14.0 //float64

**Constants**

Constants are declared like variables but with the const keyword. Constants can be character, string, boolean, or numeric values.

Note: constants cannot be declared using := syntax.

**Numeric Constants**

Numeric constants are high-precision values. An untyped constant will take the type of its value.

**Others**

%T - Shows the type of the variable
%v - Shows the value of the variable
%q - Shows strings with special characters and their escape sequences
\n - new line
