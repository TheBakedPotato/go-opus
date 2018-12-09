package opus

/*
#include <stdio.h>
#include <stdlib.h>
#include "test.h"
*/
import "C"

import "fmt"

func Random() int {
	fmt.Println(int(C.testFunction()))
	// fmt.Println(int(C.testFunction()))
	// enc := C.createEncoder()
	// opusEncoder := C.createEncoder()
	// defer C.free(unsafe.Pointer(opusEncoder))
	return int(C.random())
}
