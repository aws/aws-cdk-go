//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Connections) validateAllowDefaultPortFromParameters(other IConnectable) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowDefaultPortToParameters(other IConnectable) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowFromParameters(other IConnectable, portRange Port) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowFromAnyIpv4Parameters(portRange Port) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowInternallyParameters(portRange Port) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowToParameters(other IConnectable, portRange Port) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowToAnyIpv4Parameters(portRange Port) error {
	return nil
}

func (c *jsiiProxy_Connections) validateAllowToDefaultPortParameters(other IConnectable) error {
	return nil
}

func validateNewConnectionsParameters(props *ConnectionsProps) error {
	return nil
}

