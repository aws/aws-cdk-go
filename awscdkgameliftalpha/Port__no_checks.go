//go:build no_runtime_type_checking

package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func validatePort_TcpParameters(port *float64) error {
	return nil
}

func validatePort_TcpRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

func validatePort_UdpParameters(port *float64) error {
	return nil
}

func validatePort_UdpRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

func validateNewPortParameters(props *PortProps) error {
	return nil
}

