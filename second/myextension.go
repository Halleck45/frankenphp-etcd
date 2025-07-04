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
//#include "php-ext.h"
//#include "myextension.h"
//#include "myextension_arginfo.h"
import "C"

import (
	"fmt"
	"github.com/dunglas/frankenphp"
	"unsafe"
)

// export_php:function background_hello(string $name): void
func background_hello(foo *C.zend_string) {
	cstr := zendStringToGoString(foo)
	fmt.Println("Hello", cstr)
}

func zendStringToGoString(zendStr *C.zend_string) string {
	if zendStr == nil {
		return ""
	}

	return C.GoStringN((*C.char)(unsafe.Pointer(&zendStr.val)), C.int(zendStr.len))
}
