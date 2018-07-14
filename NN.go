package main

type NN struct {
	scores                                               [NBRROWS]float64
	nbrIntputNeurons, nbrHiddenNeurons, nbrOutputNeurons int
	nbrHiddenLayers                                      int
	learningRate                                         float64
	layers                                               []*Layer
}

func InitNN() *NN {
	nn := &NN{
		nbrHiddenLayers:  3,
		nbrIntputNeurons: 7,
		nbrHiddenNeurons: 5,
		nbrOutputNeurons: 7,
		learningRate:     0.10,
		layers:           make([]*Layer, 1+3+1),
	}

	for _, layer := range nn.layers {
		layer.InitLayer()
	}

	return nn
}
