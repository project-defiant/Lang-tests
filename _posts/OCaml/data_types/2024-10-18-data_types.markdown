---
layout: post
title:  "Data Types in OCaml"
---


OCaml is statically and strongly typed language. Types can not change
at runtime, only when explicitly converted.

We have following types:

- [x] Int (Int32 & Int64)
- [x] Float
- [x] Bool
- [x] Char (single quoted literals)
- [x] String (allows for access, immutable)
- [x] Byte Sequence 
- [x] Array of 'a type (non-extendable mutable arrays)
- [x] Lists of 'a type (extendable, non-mutable linked-lists)
- [x] Options (Some, None)
- [x] Tuples (pairs)
- [x] Functions (m -> n)
- [x] Unit (side effects)
- [x] User defined types


## User defined types

These follow the ` type … = … ` structure, where first elipsis `...` is a type name (lowercase) and second is type
definition. The definitions can be:

### Variants (tagged unions)

Enumerated types, defined by explicit list of values. Defined values are called constructors and must be capitalised.

```{OCaml}
# type rectitude = Evil | R_Neutral | Good;;
type rectitude = Evil | R_Neutral | Good
```

The variants can be recursive.

### Polymorphic types

#TODO: Add the docs here ... when you understand it !

### Records

Like dictionaries with key-value pairs (called fields).

```{OCaml}
 type character = {
  name : string;
  level : int;
  race : string;
  class_type : character_class;
  alignment : firmness * rectitude;
  armor_class : int;
};;
```

### Type aliases

Any type can have additional name.
