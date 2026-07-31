//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowSource) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (f *jsiiProxy_FlowSource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FlowSource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FlowSource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFlowSource_FromFlowSourceArnParameters(scope constructs.Construct, id *string, flowSourceArn *string) error {
	return nil
}

func validateFlowSource_FromFlowSourceAttributesParameters(scope constructs.Construct, id *string, attrs *FlowSourceAttributes) error {
	return nil
}

func validateFlowSource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFlowSource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFlowSource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFlowSourceParameters(scope constructs.Construct, id *string, props *FlowSourceProps) error {
	return nil
}

