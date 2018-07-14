package main

type NN struct {
	scores                                                       [NBRROWS]float64
	nbrIntputNeurons, nbrHiddenNeuronsPerLayer, nbrOutputNeurons int
	nbrHiddenLayers                                              int
	learningRate                                                 float64
	layers                                                       []*Layer
}

func InitNN() *NN {
	nn := &NN{
		nbrHiddenLayers:          3,
		nbrIntputNeurons:         7,
		nbrHiddenNeuronsPerLayer: 5,
		nbrOutputNeurons:         7,
		learningRate:             0.10,
	}
	nn.layers = make([]*Layer, 1+nn.nbrHiddenLayers+1)

	nn.layers[0].InitLayer(nn.nbrIntputNeurons)
	for numLayer := 1; numLayer < len(nn.layers)-1; numLayer++ {
		nn.layers[numLayer].InitLayer(nn.nbrHiddenNeuronsPerLayer)
	}
	nn.layers[len(nn.layers)].InitLayer(nn.nbrOutputNeurons)

	return nn
}
