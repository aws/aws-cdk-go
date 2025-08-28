//go:build no_runtime_type_checking

package awselasticloadbalancingv2targets

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerTarget) validateAttachToNetworkTargetGroupParameters(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) error {
	return nil
}

func validateNewAlbListenerTargetParameters(albListener awselasticloadbalancingv2.ApplicationListener) error {
	return nil
}

