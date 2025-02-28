//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VirtualRouter) validateAddRouteParameters(id *string, props *RouteBaseProps) error {
	return nil
}

func (v *jsiiProxy_VirtualRouter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VirtualRouter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VirtualRouter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVirtualRouter_FromVirtualRouterArnParameters(scope constructs.Construct, id *string, virtualRouterArn *string) error {
	return nil
}

func validateVirtualRouter_FromVirtualRouterAttributesParameters(scope constructs.Construct, id *string, attrs *VirtualRouterAttributes) error {
	return nil
}

func validateVirtualRouter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVirtualRouter_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVirtualRouter_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVirtualRouterParameters(scope constructs.Construct, id *string, props *VirtualRouterProps) error {
	return nil
}

