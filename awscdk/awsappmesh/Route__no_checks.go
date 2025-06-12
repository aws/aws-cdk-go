//go:build no_runtime_type_checking

package awsappmesh

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Route) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Route) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateRoute_FromRouteArnParameters(scope constructs.Construct, id *string, routeArn *string) error {
	return nil
}

func validateRoute_FromRouteAttributesParameters(scope constructs.Construct, id *string, attrs *RouteAttributes) error {
	return nil
}

func validateRoute_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRoute_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRoute_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRouteParameters(scope constructs.Construct, id *string, props *RouteProps) error {
	return nil
}

