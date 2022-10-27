//go:build no_runtime_type_checking

package awselasticloadbalancingv2targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func validateNewAlbTargetParameters(alb awselasticloadbalancingv2.ApplicationLoadBalancer, port *float64) error {
	return nil
}

