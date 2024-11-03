let x = 
  match not true with
    | true -> "yep"
    | false -> "oops";;

let y = 
  match 42 with
  | foo -> foo;;


let z = 
  match "foo" with
    | "bar" -> 0
    | "baz" -> 0
    | "foo" -> 1
    | _ -> 0;;



