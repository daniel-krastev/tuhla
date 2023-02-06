The `BUILD` file in this dir uses `go_module`, which requires manual dependencies declaration. The 
new `go/BUILD` uses `go_get`, which only needs the modules to download as found in `go.mod`. The
folder is currently not used but kept for reference.
