//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpcEndpointServiceDomainName) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (v *jsiiProxy_VpcEndpointServiceDomainName) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateVpcEndpointServiceDomainName_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_VpcEndpointServiceDomainName) validateSetDomainNameParameters(val *string) error {
	return nil
}

func validateNewVpcEndpointServiceDomainNameParameters(scope constructs.Construct, id *string, props *VpcEndpointServiceDomainNameProps) error {
	return nil
}

