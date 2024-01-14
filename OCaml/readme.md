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

### Lists

lists are created with following syntax:

```{OCaml}
let empty_list = [];;
let int_list = [1;2;3];;
```

Important note is that list elements are split by `semicolon`, not comma.

