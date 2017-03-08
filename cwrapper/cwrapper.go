package cwrapper

/*

#cgo LDFLAGS: -L${SRCDIR} -lmyc

#include <stdio.h>
#include "myc.h"

void println(char* str)
{
    printf("%s\n", str);
}

*/
import "C"

func Println(s string) {
    str := C.CString(s);
    C.println(str);
}

func MyFoo() {
    C.my_foo();
}

