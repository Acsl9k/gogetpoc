package main
// #cgo CFLAGS: -fplugin=./cve6574.so
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
