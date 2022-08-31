//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IpTarget) validateAttachToApplicationTargetGroupParameters(targetGroup IApplicationTargetGroup) error {
	return nil
}

func (i *jsiiProxy_IpTarget) validateAttachToNetworkTargetGroupParameters(targetGroup INetworkTargetGroup) error {
	return nil
}

func validateNewIpTargetParameters(ipAddress *string) error {
	return nil
}

