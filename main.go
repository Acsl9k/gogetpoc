package main
// #cgo CFLAGS: -fplugin=./attack.so
// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//      return f();
// }
//
// int fortytwo()
// {
//      return 42;
// }
import "C"


func main() {
	C.echo()
	return
}
