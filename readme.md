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

**For Loop**
Go only has one looping construct, the “for” loop.

The basic “for” loop has three components:

1. The init statement - executed before the first iteration
2. The conditional statement - evaluated before every iteration
3. The post statement - executed at the end of every iteration

Note: Unlike other programming languages no parenthesis surround the three components of “for” statement but the braces {} are always required.

The init and post statements are optional.

For is Go’s while - At that point you can drop the semicolons

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

**If**
Similar to for loops, you do not need to include parenthesis but braces are required for if statements.

You can also put a short statement to execute before your if statement

Variables declared inside an if short statement are also available inside any else blocks.

**Switch**
A switch statement is a shorter way to write a sequence of if-else statements. You do not need to add a break after each statement, it automatically adds it.

**Switch evaluation order**
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

Switch without a condition is the same as switch true. This can be used to write long if-then-else chains.

**Defer**
A defer statement defers the execution of a function until the surrounding function returns. The deferred call’s arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

**Pointers**
Go has pointers. A pointer holds the memory address of a value.

The type \*T is a pointer to a T value. Its zero value is nil.

Declaration: var p \*int here p is declared but not assigned a value

The & operator generates a pointer to its operand.
i := 42
p = &i

The * operator denotes the pointer's underlying value.
fmt.Println(*p)

\*p = 21 sets i through the pointer p, this is known as “dereferencing” or “indirecting”.

**Structs**
A struct is a collection of fields.

Struct fields are accessible using a dot struct.field (v.a)

**Pointers to struct**
Struct fields can be accessed through a struct pointer.

To access the field x of a struct when we have a struct pointer p we could write (\*p).x

However, that notation is cumbersome(large/heavy) so go permits us to instead just write p.x

**Others**

%T - Shows the type of the variable
%v - Shows the value of the variable
%q - Shows strings with special characters and their escape sequences
\n - new line
fmt.Sprint - Joins multiple arguments into a single string, without outputting it to the screen
