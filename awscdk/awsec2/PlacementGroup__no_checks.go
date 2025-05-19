//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PlacementGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PlacementGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PlacementGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePlacementGroup_FromPlacementGroupNameParameters(scope constructs.Construct, id *string, placementGroupName *string) error {
	return nil
}

func validatePlacementGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePlacementGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePlacementGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPlacementGroupParameters(scope constructs.Construct, id *string, props *PlacementGroupProps) error {
	return nil
}

