package main

/*
#include <stdlib.h>
#include "main.h"
*/
import "C"
import "runtime/cgo"
import "fmt"

func init() {
	frankenphp.RegisterExtension(unsafe.Pointer(&C.ext_module_entry))
}


//export background_hello
func BackgroundHello(foo *C.zend_string) {
	cstr := C.zend_string(foo)
	fmt.Println("Hello", C.zend_string(cstr))
}

