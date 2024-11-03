(* Extend list 1 with list 2 
The issue with OCaml lists is that they store only a pointer to
the left most element
*)
let rec extend lst1 lst2 = 
  match lst1 with
    | [] -> lst2
    | head :: tail -> head :: extend tail lst2


(* Calculate lenght of the list *)
let rec length lst = 
  match lst with 
    | [] -> 0
    | head :: tail -> 1 + length tail


(* Calculate sum of elements of list[int] *)
let rec sum lst = 
  match lst with
  | [] -> 0
  | head :: tail -> head + sum tail


(* Check if list is empty*)
let empty lst = 
  match lst with
   | [] -> true
   | _ :: _ -> false


(* By default the match is performed onto the last argument, so this holds*)

let empty2 = function
  | [] -> true
  | _ :: _ -> false
