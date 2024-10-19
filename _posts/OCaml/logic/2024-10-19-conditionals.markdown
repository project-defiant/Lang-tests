---

layout: post
title:  "OCaml conditionals"

---

Below are some code snippets on how to use logical expressions in OCaml.

All examples are run in `utop`

```ocaml
if "batman" > "superman"; then "yay" else "boo";;
(* - : string = "boo" *)
```

OCaml is strict about the type checking, so the evaluation as below will not work!

```ocaml
if 0; then "yay" else "boo";;
(* Error: This expression has type int but an expression was expected of type bool because it is in the condition of an if-statement *)

if "batman" > "superman"; then "yay";;
(* Error: This expression has type string but an expression was expected of type unit because it is in the result of a conditional with no else branch *)

```

The expression has following properties


`if expression_1: bool; then expression_2: T else expression_3 : T`

This get's evaluated based on `expression_1` outcome, that has to be boolean value. Based it evaluates to `true` the `expression_2` get's evaluated. Otherwise the `expression_3` get's evaluated. The result is either `expression_2`, `expression_3` or error.
