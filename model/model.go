package model

func CreateNetworkModel(layersSize []int) *Layer {
	argCount := len(layersSize)
	if argCount < 2 {
		panic("MINIMAL LAYER SIZE IS 2")
	}

	layer1 := NewLayer(layersSize[0])

	for i := 1; i < argCount; i++ {
		size := layersSize[i]
		layer2 := NewLayer(size)
		layer1.ConnectLayer(layer2)
	}

	return layer1
}
