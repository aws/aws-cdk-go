//go:build no_runtime_type_checking

package awscdklocationalpha

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Map) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_Map) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_Map) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (m *jsiiProxy_Map) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (m *jsiiProxy_Map) validateGrantRenderingParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateMap_FromMapArnParameters(scope constructs.Construct, id *string, mapArn *string) error {
	return nil
}

func validateMap_FromMapNameParameters(scope constructs.Construct, id *string, mapName *string) error {
	return nil
}

func validateMap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMap_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMap_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewMapParameters(scope constructs.Construct, id *string, props *MapProps) error {
	return nil
}

