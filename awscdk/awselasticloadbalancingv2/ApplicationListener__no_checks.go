//go:build no_runtime_type_checking

package awselasticloadbalancingv2

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApplicationListener) validateAddActionParameters(id *string, props *AddApplicationActionProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateAddCertificatesParameters(id *string, certificates *[]IListenerCertificate) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateAddTargetGroupsParameters(id *string, props *AddApplicationTargetGroupsProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateAddTargetsParameters(id *string, props *AddApplicationTargetsProps) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_ApplicationListener) validateRegisterConnectableParameters(connectable awsec2.IConnectable, portRange awsec2.Port) error {
	return nil
}

func validateApplicationListener_FromApplicationListenerAttributesParameters(scope constructs.Construct, id *string, attrs *ApplicationListenerAttributes) error {
	return nil
}

func validateApplicationListener_FromLookupParameters(scope constructs.Construct, id *string, options *ApplicationListenerLookupOptions) error {
	return nil
}

func validateApplicationListener_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplicationListener_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateApplicationListener_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewApplicationListenerParameters(scope constructs.Construct, id *string, props *ApplicationListenerProps) error {
	return nil
}

