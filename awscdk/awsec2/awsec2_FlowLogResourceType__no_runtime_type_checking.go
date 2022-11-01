//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateFlowLogResourceType_FromNetworkInterfaceIdParameters(id *string) error {
	return nil
}

func validateFlowLogResourceType_FromSubnetParameters(subnet ISubnet) error {
	return nil
}

func validateFlowLogResourceType_FromVpcParameters(vpc IVpc) error {
	return nil
}

func (j *jsiiProxy_FlowLogResourceType) validateSetResourceIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FlowLogResourceType) validateSetResourceTypeParameters(val *string) error {
	return nil
}

