package model

type Connect struct {
	Input  Input
	Widget float64
}

// math.SmallestNonzeroFloat64
type Input struct {
	Value    float64
	Connects []Connect
}

func (input Input) ConnectIn(connectMe Input) Connect {
	connect := Connect{
		Input:  connectMe,
		Widget: 0,
	}
	input.Connects = append(input.Connects, connect)
	return connect
}

func (input Input) ConnectOut(connectMe Input) Connect {
	connect := Connect{
		Input:  input,
		Widget: 0,
	}
	connectMe.Connects = append(connectMe.Connects, connect)
	return connect
}

func (input Input) TrainMe(connectMe Input) Connect {
	connect := Connect{
		Input:  input,
		Widget: 0,
	}
	connectMe.Connects = append(connectMe.Connects, connect)
	return connect
}
