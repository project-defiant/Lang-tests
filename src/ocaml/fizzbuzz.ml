(* Not so simple fizzbuzz*)

let getMsg (v: int) : string = 
  let v = v in 
    let elligible = ( v mod 3 == 0, v mod 5 == 0) in 
      match elligible with
        | (true, true) -> "FizzBuzz"
        | (true, false) -> "Fizz"
        | (false, true) -> "Buzz"
        | (false, false) -> (let x = string_of_int v in x)
  

let fizzbuzz (v: int) : unit =
  let lst = List.init v (fun x -> x + 1) in
    let rec loop f l = 
      match l with
        | [] -> ()
        | x :: xs -> print_endline (f x); loop f xs in
      loop getMsg lst