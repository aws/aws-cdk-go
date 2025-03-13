//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplicationListener) validateAddActionParameters(id *string, props *AddApplicationActionProps) error {
	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddCertificatesParameters(id *string, certificates *[]IListenerCertificate) error {
	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddTargetGroupsParameters(id *string, props *AddApplicationTargetGroupsProps) error {
	return nil
}

func (i *jsiiProxy_IApplicationListener) validateAddTargetsParameters(id *string, props *AddApplicationTargetsProps) error {
	return nil
}

func (i *jsiiProxy_IApplicationListener) validateRegisterConnectableParameters(connectable awsec2.IConnectable, portRange awsec2.Port) error {
	return nil
}

func (i *jsiiProxy_IApplicationListener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

