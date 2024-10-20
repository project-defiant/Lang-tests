---
layout: post
title:  "Interfaces"
---

> The main idea behind the interface is that it implements a function signature, but not the concrete behavior.

After some time spent in python and R I have a strong opinion of what nice in both languages OOP design.
From python I love the way to define the classes and how easy is to define and use them with a single `__init__` method.

Also the `@classmethod`, `@staticmethod` and `@property` are everything that a goog OOP language should have.


There is only one catch with python OOP design (and the OOP design in Java and other languages) - **Inheritance**.
The inheritance can make your programm way more difficult to track. I am sure if you are familiar with the java language, 
you know how long the stack traces can be there. Besides that you have to make clean representation and semantically bind the
class entities (methods and properties) in a way to make sure that the entity should actually belong to the current class, rather 
then to it's predecessor.

In a way this is good, but only when the real world problems can be encapsulated by extracting common derivative, as most likely during the
code evolution you will eventually need to update the methods of the parent classes, as more and more entities will be common to many different classes that use the same inheritance.

One solution to this dilema is to not use any inheritance if it comes to the common programming problems like:

* print an object
* iterate over the object
* cast the object
* hash the object

and anything else that is a common action that could be coded by a programmer. 

These kind of `actions` (as I would call them) are very `generic` in terms of the outcome, so they share common return value.
Their inputs may vary though, depending on the object that the action need's to be performed on. For instance in terms of calculating
the simple `area` of the figure, when it comes to the circle, you will need the radius, on the other hand with rectangle you will need the 
lenght of the two perpendicular axes. 

Given that properties we could construct a signature of an action as follows:

```
action area(object: Any) -> (area: Integer)
```

For the `Shape` class we could create a method `area` and not implement it. This is a common practice in python. Then inherit the `Shape` over in the `Rectangle` class like this:


```python
class Shape:
    def area() -> int:
      raise NotImplementedError("Implement in child classes")

class Rectangle(Shape):
    def __init__(self, width: int, height: int) -> None:
        self.width = width
        self.height = height

    def area(self) -> int:
        return self.width * self.height
```

The above will be the simplest implementation of an **Interface** in python.


Now python does not allow `strictly` speaking for the interfaces, but you could substitute it with the `Protocol` class that implements
common behavior for python class methods that satisfy it. This is the case, because python would not make of any use of the interfaces during
the writing of the code, as python does compile the program just before it gets executed. 

On the other hand many programming languages already introduced the `interfaces` in one way or the other. The examples can be found in GO, Rust, Scala or even Java.


Thought how it looks, there is one caviat to the fact of abusing the inheritance inside the interfaces, as to my strong opinion they should be actually separate things. Here comes R with it's **very old** S3 objects. This OOP method although, not appreciated by many people, to me
it's best written implementation of iterfaces. I encourage you to read more about it, but for the sake of this article I will just explain how the S3 works and why you can print (almost everything) in R and even python.

S3 is a design decision, that allows you to implement a simple object (without the synctactic sugar) in R. This comes along with the concept of `generics` - methods that are meaningful for most objects that can be. With the S3 you can design a simple object and assign it a name of a class. You do not implement any specific methods over the class though! you implement them outside of the class. The way why the S3 is unique with regards to python implementation of generics with `__dunder__methods__` is that R does not couple the exact method with the object itself but rather creates a namespace that links to the object if needed. This way we do not need to satisfy the concrete implementation (if we do not intent do do so) even if the parent class would require it as in the python example above.


Thought one, as is, but I consider the interfaces far more superor when it comes to defining common methods, then implementing them 
inside the parent classes.

May the interfaces be with us!





