#ifndef _MAIN_H
#define _MAIN_H

#include <php.h>
#include <stdint.h>

extern zend_module_entry ext_module_entry;

typedef struct go_value go_value;

typedef struct go_string {
  size_t len;
  char *data;
} go_string;



#endif
