//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VdmAttributes) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VdmAttributes) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VdmAttributes) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateVdmAttributes_FromVdmAttributesNameParameters(scope constructs.Construct, id *string, vdmAttributesName *string) error {
	return nil
}

func validateVdmAttributes_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVdmAttributes_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVdmAttributes_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVdmAttributesParameters(scope constructs.Construct, id *string, props *VdmAttributesProps) error {
	return nil
}

