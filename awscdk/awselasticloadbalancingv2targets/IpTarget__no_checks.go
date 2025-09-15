//go:build no_runtime_type_checking

package awselasticloadbalancingv2targets

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpTarget) validateAttachToApplicationTargetGroupParameters(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) error {
	return nil
}

func (i *jsiiProxy_IpTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func validateNewIpTargetParameters(ipAddress *string) error {
	return nil
}

