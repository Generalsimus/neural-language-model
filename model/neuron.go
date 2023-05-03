package model

type ConnectN struct {
	Neuron Neuron
	Widget float64
}

type Neuron struct {
	Value    float64
	Connects []ConnectN
}

// გარემოს აღქმა,ნივთების მსგავსების აღქმა 2d აღქმა
func (neuron Neuron) ConnectIn(connectMe Neuron) ConnectN {
	connect := ConnectN{
		Neuron: connectMe,
		Widget: 0,
	}
	neuron.Connects = append(neuron.Connects, connect)
	return connect
}

func (neuron Neuron) ConnectOut(connectMe Neuron) ConnectN {
	connect := ConnectN{
		Neuron: neuron,
		Widget: 0,
	}
	connectMe.Connects = append(connectMe.Connects, connect)
	return connect
}
