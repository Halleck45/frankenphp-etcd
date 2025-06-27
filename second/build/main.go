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
	cstr := C.get_zend_string_val(foo)
	fmt.Println("Hello", C.GoString(cstr))
}

