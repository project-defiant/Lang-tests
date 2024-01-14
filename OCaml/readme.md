# OCaml language


## main concepts

### Variables

OCaml is a function programming language. It uses following syntax.

```{OCaml}
let x = 50;;
```

Variables are immutable (they are constant). To use mutable variables one need to use `references`.
Variables can be limited to some expression scope only
```{OCaml}
let y = 50 in y * y;;
```
The code above returns `2500` without assigning the result to the variable `y`.
Variable `y` is only present in the scope followed by `in` expression.

If one want to assign the result of `in` expression, then they can do following:

```{OCaml}
let s = let y = 50 in let z = 10 in y * z;;
```

Variable `s` is initialized with value of `500`. `y` and `z` are only bound to the expression scope.
### Lists

lists are created with following syntax:

```{OCaml}
let empty_list = [];;
let int_list = [1;2;3];;
```

Important note is that list elements are split by `semicolon`, not comma.

To prepend to the list use `cons` operator. It is not symmetric, so appending does not work.
```{OCaml}
let int_list = [1;2;3];;
9 :: int_list;;
```

### references

To create reference use `ref` keyword.
```{OCaml}
# let r = ref 0;;
val r : int ref = {contents = 0}
```
Accessing the reference needs to be done with `!` dereference operator.
```{OCaml}
# !r;;
- : int = 0
```
To update the reference content use `:=`

```{OCaml}
# r:= 10;;
- : int = 10
```


One can chain couple of expressions with a single `;` between them.

### functions

Functions are defined as follow:

```{OCaml}
let square x = x * x;;
square 10 ;;
square(20);;
```

Anonymous functions

```{OCaml}
fun x -> x * x;;
let s = (fun x -> x * x) 50;;
```


Partial execution of a function


```{OCaml}
let cat x y -> x ^ " " ^ y;;
let partial = cat "Hello";;
let hello_world = partial "World";;
```
If one do not pass all arguments to the function. It becomes partial, and expects other arguments to 
fill the blank ones before it is executed.

Side effects are performed by set of functions like:
- [x] read_line
- [x] print_endline
the first one takes IO (unit type), the latter returns to IO (unit type)


Recursive functions should be used instead of `for` and `while` loops.

```{OCaml}
let rec range lo hi = 
  if lo > hi then
    []
  else 
    lo :: range (lo + 1) hi;;
```

With recursive function they have to start with `let rec _ = _`



## Pattern matching

We can match anything except functions.

```{OCaml}
let g' x = match x with
    | "foo" -> 1
    | "bar" -> 2
    | "baz" -> 3
    | "qux" -> 4
    | _ -> 0;;
val g' : string -> int = <fun>
```


## Tuples

Tuples can have mixed types.

```{OCaml}
let tpl = (1,2,"aaa");;
```
To access tuple element use pattern matching.

```{OCaml}
let second_elem t = match t with
  | (_, y, _) -> y;;
```

## types

### Variant types

One can use enum like types.
```{OCaml}
# type page_range =
    | All
    | Current
    | Range of int * int;;
type page_range = All | Current | Range of int * int

```

### Record types

The record types are a bit like hash maps, one can access value by dot.

```{OCaml}
# type person = {
    first_name : string;
    surname : string;
    age : int
  };;
type person = { first_name : string; surname : string; age : int; }
```

## Exception

To raise an exception one should use `raise`. They are caught by `try with` expression.

```{OCaml}
# let id_42 n = if n <> 42 then raise (Failure "Sorry") else n;;
val id_42 : int -> int = <fun>

# id_42 42;;
- : int = 42

# id_42 0;;
Exception: Failure "Sorry".

# try id_42 0 with Failure _ -> 0;;
- : int = 0
```

## Modules

One can see what is in the module with `#show {Module}`

To call module function one need to access it by it's module name:

```{OCaml}
Option.map (fun x -> x * x) None;;
Option.map (fun x -> x * x) (Some 8);;
```

