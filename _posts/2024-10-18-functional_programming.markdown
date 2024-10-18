---
layout: post
title:  "Notes from functional programming"
---


**disclaimer - notes comes from [Michael Ryan Clarkson Open Course](https://www.cs.cornell.edu/~clarkson/)

The textbook is "OCaml Programming: Correct + Efficient + Beautiful": https://cs3110.github.io/textbook

* defines computations as mathematical functions
* avoids mutable state

Expressions specify what to compute

* Variables never change value
* Functions never have side effects

## Imperative program

* Change state
Examples of commands that change the state:

```python
x = x + 1
a[i] = 52
p.next = p.next.next
```

* Side effects
  
```c
int x= 0
int incr_x() {
  x++;
  return x;
}
```
