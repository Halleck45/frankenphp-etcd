package myextension

/*
#cgo CFLAGS: -Wall
#include "php.h"
#include "build/main.h"
#include "build/main_arginfo.h"
#cgo darwin pkg-config: libxml-2.0
#cgo CFLAGS: -Wall -Werror
#cgo CFLAGS: -I/usr/local/include -I/usr/local/include/php -I/usr/local/include/php/main -I/usr/local/include/php/TSRM -I/usr/local/include/php/Zend -I/usr/local/include/php/ext -I/usr/local/include/php/ext/date/lib
#cgo linux CFLAGS: -D_GNU_SOURCE
#cgo darwin CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/usr/local/lib -L/usr/lib -lphp -lm -lutil
#cgo linux LDFLAGS: -ldl -lresolv
#cgo darwin LDFLAGS: -Wl,-rpath,/usr/local/lib -L/opt/homebrew/lib -L/opt/homebrew/opt/libiconv/lib -liconv -ldl
#include <stdlib.h>
#include <stdint.h>
#include <php_variables.h>
#include <zend_llist.h>
#include <zend_string.h>

// Wrapper to access zend_string value
static const char* get_zend_string_val(zend_string* str) {
    return ZSTR_VAL(str);
}
*/
import "C"

import (
	"fmt"
)

// export_php:function background_hello(string $name): void
func BackgroundHello(foo *C.zend_string) {
	cstr := C.get_zend_string_val(foo)
	fmt.Println("Hello", C.GoString(cstr))
}
