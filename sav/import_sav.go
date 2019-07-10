package sav

// #cgo amd64 CFLAGS: -g
// #cgo LDFLAGS: -lreadstat
// #include <stdlib.h>
// #include "import_sav.h"
import "C"

import (
	"fmt"
	"unsafe"
)

type headerLine struct {
	name, vType string
}

var headerItems = make(map[int]headerLine)

//export goAddLine
func goAddLine(str *C.char) {
	//gostr := C.GoString(str)
	//println(gostr)
}

//export goAddHeaderLine
func goAddHeaderLine(pos C.int, name *C.char, varType *C.char, end C.int) {
	if int(end) == 1 { // we are done
		for k := range headerItems {
			fmt.Printf("key[%d] -> title[%s] -> type[%s]\n", k, headerItems[k].name, headerItems[k].vType)
		}
	} else {
		headerItems[int(pos)] = headerLine{C.GoString(name), C.GoString(varType)}
	}
}

func Import(fileName string) int {
	name := C.CString(fileName)
	defer C.free(unsafe.Pointer(name))

	res := C.parse_sav(name)
	if res != 0 {
		return 1
	}

	return 0
}
