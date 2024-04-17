package modelLoader

// #cgo LDFLAGS: -L/home/ayush.goya/minimalExample/modelLoader/ -l_lightgbm
// #include "c_api.h"
// #include <stdio.h>
// #include <stdlib.h>
import "C"

var predictor C.BoosterHandle

func Load(modelString string) {
	outNumIterations := C.int(0)
	var modelPtr *C.char = C._GoStringPtr(modelString)
	res := int(C.LGBM_BoosterLoadModelFromString(modelPtr, &outNumIterations, &predictor))
	if res == -1 {
		panic("LightGBM C_API : Failed to load model from the model string")
	}
	//println("Load Success")
}

func ReleaseMemory() {
	res := int(C.LGBM_BoosterFree(predictor))

	if res == -1 {
		panic("LightGBM C_API : Failed to release the memory of the LightGBM model")
	}
	//println("ReleaseMemory Success")
}
