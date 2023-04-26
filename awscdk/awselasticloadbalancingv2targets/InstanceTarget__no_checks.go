//go:build no_runtime_type_checking

package awselasticloadbalancingv2targets

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceTarget) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (i *jsiiProxy_InstanceTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func validateNewInstanceTargetParameters(instance awsec2.Instance) error {
	return nil
}

