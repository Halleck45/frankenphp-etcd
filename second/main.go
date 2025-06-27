package myextension

//#include <php_config.h>
//#include "php.h"
//#include "zend.h"
//#include <zend_exceptions.h>
//#include <stdlib.h>
//#include <stdint.h>
//#include <php_variables.h>
//#include <zend_llist.h>
//#include <zend_string.h>
//#include "main.h"
//#include "main_arginfo.h"
import "C"

import (
	"fmt"
)

// export_php:function background_hello(string $name): void
func BackgroundHello(foo *C.zend_string) {
	cstr := C.zend_string(foo)
	fmt.Println("Hello", C.zend_string(cstr))
}
