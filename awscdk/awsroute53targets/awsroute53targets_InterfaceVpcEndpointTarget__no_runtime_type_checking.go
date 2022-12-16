//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InterfaceVpcEndpointTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewInterfaceVpcEndpointTargetParameters(vpcEndpoint awsec2.InterfaceVpcEndpoint) error {
	return nil
}

