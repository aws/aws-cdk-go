//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DsRecord) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DsRecord) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DsRecord) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateDsRecord_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDsRecord_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDsRecord_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDsRecordParameters(scope constructs.Construct, id *string, props *DsRecordProps) error {
	return nil
}

