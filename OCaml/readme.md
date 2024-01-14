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

