let print_ranges ranges =
  let parts = List.map (fun (a, b) -> Printf.sprintf "%d-%d" a b) ranges in
  print_endline (String.concat " " parts)

let cafeteria () =
  let rec create_ranges (acc : (int * int) list) =
    try
      let line = read_line () in
      let range_list =
        String.split_on_char '-' line |> List.map int_of_string
      in
      let tuple = match range_list with [ a; b ] -> (a, b) in
      create_ranges (tuple :: acc)
    with _ -> acc
  in
  let ranges = create_ranges [] |> List.sort compare in
  (* print_ranges ranges; *)
  let rec loop current acc = function
    | [] -> List.rev (current :: acc)
    | (a, b) :: tl ->
        let ca, cb = current in
        if a <= cb + 1 then loop (ca, max cb b) acc tl
        else loop (a, b) (current :: acc) tl
  in
  let fusion = match ranges with [] -> [] | h :: t -> loop h [] t in
  (* print_ranges fusion; *)
  fusion |> List.fold_left (fun s (a, b) -> s + (b - a + 1)) 0
;;

print_endline (string_of_int (cafeteria ()))