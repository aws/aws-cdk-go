//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Resource) validateApplyRemovalPolicyParameters(policy RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Resource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string, props *ResourceProps) error {
	return nil
}

