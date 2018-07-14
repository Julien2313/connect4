package main

type Layer struct {
	neurals []*Neural
}

func (l *Layer) InitLayer(nbrNeurals int) {
	l.neurals = make([]*Neural, nbrNeurals)

}
