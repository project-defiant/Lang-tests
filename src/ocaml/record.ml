(* Record type *)

type carthesianPoint = {
  x : float;
  y : float;
  z : float;
}

type sphericalPoint = {
  r : float;
  lat: float;
  long: float;
}


(* Create an instance of the carthesian point*)
let p1: carthesianPoint = {
  x=1.;
  y=2.;
  z=3.;  
}

(* Convert spherical point to carthesian in 3d *)
let toSpherical (p: carthesianPoint): sphericalPoint =
  let r = sqrt (p.x *. p.x +. p.y *. p.y +. p.z *. p.z) in
  let long = acos(p.x /. sqrt(p.x *. p.x +. p.y *. p.y)) in
  let lat = acos(p.z /. r) in
  { r; lat; long; }

(* Convert carthesian point to spherical point in 3d *)
let toCarthesian (p: sphericalPoint) : carthesianPoint =
  let x = p.r *. cos p.lat *. cos p.long in
  let y = p.r *. cos p.lat *. sin p.long in
  let z = p.r *. sin p.lat in
  {x; y; z}
