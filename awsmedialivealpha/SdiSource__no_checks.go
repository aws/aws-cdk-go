//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SdiSource) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (s *jsiiProxy_SdiSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SdiSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SdiSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSdiSource_FromSdiSourceAttributesParameters(scope constructs.Construct, id *string, attrs *SdiSourceAttributes) error {
	return nil
}

func validateSdiSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSdiSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSdiSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSdiSourceParameters(scope constructs.Construct, id *string, props *SdiSourceProps) error {
	return nil
}

