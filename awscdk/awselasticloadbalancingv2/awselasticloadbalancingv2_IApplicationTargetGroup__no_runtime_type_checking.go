//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplicationTargetGroup) validateRegisterConnectableParameters(connectable awsec2.IConnectable) error {
	return nil
}

func (i *jsiiProxy_IApplicationTargetGroup) validateRegisterListenerParameters(listener IApplicationListener) error {
	return nil
}

