# Values and functions in OCaml

Everything is an expression that yields a value. OCaml type system has a property of `subject reduction`.
This property tells us that the type will not change at runtime (for instance no implicit type conversions).

Each expression can be named, and evaluated or not, depending on the expression value.

For instance:

```{OCaml}
utop # let value = "something";;
val value : string = "something"
utop # let value2 x  = x * x;;
val value2 : int -> int = <fun>
```
## Local and top level expressions

One can have global (top level) or local (using `in` keyword) named expressions. Local definitions can be
- [x] chained

```{OCaml}
# let d = 2 * 3 in
  let e = d * 7 in
  d * e;;
- : int = 252

# d;;
Error: Unbound value d
# e;;
Error: Unbound value e
```
- [x] nested

```{OCaml}
# let d =
    let e = 2 * 3 in
    e * 5 in
  d * 7;;
- : int = 210

# d;;
Error: Unbound value d
# e;;
Error: Unbound value e
```

## Pattern matching and destructuring

One can create a descructuring assingnment with `tuples` and `records`.

## functions

In OCaml one do not call functions with parenthesis, rather by it's name and pass all positional arguments afterwards.
The parenthesis are used to limit the expressions to subexpressions only.

Reverse pipe operator in OCaml is called `application operator` and it's value is `@@`.

Original pipe operator is `|>`.
```{OCaml}
# int_of_float (sqrt (float_of_int (int_of_string "81")));;
- : int = 9

# int_of_float @@ sqrt @@ float_of_int @@ int_of_string "81";;
- : int = 9
# "81" |> int_of_string |> float_of_int |> sqrt |> int_of_float;;
- : int = 9
```



