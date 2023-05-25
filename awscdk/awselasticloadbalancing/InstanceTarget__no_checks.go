//go:build no_runtime_type_checking

package awselasticloadbalancing

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceTarget) validateAttachToClassicLBParameters(loadBalancer LoadBalancer) error {
	return nil
}

func validateNewInstanceTargetParameters(instance awsec2.Instance) error {
	return nil
}

