//go:build no_runtime_type_checking

package awsefs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessPoint) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AccessPoint) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AccessPoint) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAccessPoint_FromAccessPointAttributesParameters(scope constructs.Construct, id *string, attrs *AccessPointAttributes) error {
	return nil
}

func validateAccessPoint_FromAccessPointIdParameters(scope constructs.Construct, id *string, accessPointId *string) error {
	return nil
}

func validateAccessPoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccessPoint_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAccessPoint_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAccessPointParameters(scope constructs.Construct, id *string, props *AccessPointProps) error {
	return nil
}

