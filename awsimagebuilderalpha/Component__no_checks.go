//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Component) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Component) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Component) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_Component) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (c *jsiiProxy_Component) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateComponent_FromComponentArnParameters(scope constructs.Construct, id *string, componentArn *string) error {
	return nil
}

func validateComponent_FromComponentAttributesParameters(scope constructs.Construct, id *string, attrs *ComponentAttributes) error {
	return nil
}

func validateComponent_FromComponentNameParameters(scope constructs.Construct, id *string, componentName *string) error {
	return nil
}

func validateComponent_IsComponentParameters(x interface{}) error {
	return nil
}

func validateComponent_IsConstructParameters(x interface{}) error {
	return nil
}

func validateComponent_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateComponent_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewComponentParameters(scope constructs.Construct, id *string, props *ComponentProps) error {
	return nil
}

