---
layout: post
title:  "Values and functions in OCaml"
---

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

All recursive functions must be bound to a name.

If we use `fun` keyword, we must use the `->` symbol. 
We can use closures.

The functions always have a single parameter (if multiple parameters are applied, then we use partials).
We can also use tuple parameters in single parameter functions. 

1. Function that uses multiple arguments is known as curried `string -> (string -> string)`
2. Function that uses tuple as argument is known as uncurried `string * string -> string`

Conversion `1 -> 2` is called `uncurrying`, the other way around is `currying`.

Another important concepts are `domain` and `codomain`.
`domain (inputs) -> codomain (results)`. If the function operates outside their domain or codomain.
this means that is performs some `side effects`. `Pure` functions do not have side effects. `Impure` functions
do.
