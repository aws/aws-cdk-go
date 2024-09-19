//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RouteTable) validateAddRouteParameters(id *string, destination *string, target RouteTargetType) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RouteTable) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRouteTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRouteTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRouteTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRouteTableParameters(scope constructs.Construct, id *string, props *RouteTableProps) error {
	return nil
}

