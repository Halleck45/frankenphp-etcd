#include <php.h>
#include <Zend/zend_API.h>
#include <stddef.h>

#include "main.h"
#include "main_arginfo.h"
#include "_cgo_export.h"


PHP_MINIT_FUNCTION(main) {
    
    return SUCCESS;
}

zend_module_entry main_module_entry = {STANDARD_MODULE_HEADER,
                                         "main",
                                         ext_functions,             /* Functions */
                                         PHP_MINIT(main),  /* MINIT */
                                         NULL,                      /* MSHUTDOWN */
                                         NULL,                      /* RINIT */
                                         NULL,                      /* RSHUTDOWN */
                                         NULL,                      /* MINFO */
                                         "1.0.0",                   /* Version */
                                         STANDARD_MODULE_PROPERTIES};

PHP_FUNCTION(background_hello)
{
    zend_string *name = NULL;
    ZEND_PARSE_PARAMETERS_START(1, 1)
        Z_PARAM_STR(name)
    ZEND_PARSE_PARAMETERS_END();
    background_hello(name);
}

