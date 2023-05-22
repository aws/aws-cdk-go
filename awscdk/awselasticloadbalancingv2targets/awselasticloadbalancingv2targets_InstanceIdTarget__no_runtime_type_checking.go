//go:build no_runtime_type_checking

package awselasticloadbalancingv2targets

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InstanceIdTarget) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (i *jsiiProxy_InstanceIdTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func validateNewInstanceIdTargetParameters(instanceId *string) error {
	return nil
}

