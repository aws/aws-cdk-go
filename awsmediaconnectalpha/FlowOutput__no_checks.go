//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowOutput) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (f *jsiiProxy_FlowOutput) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FlowOutput) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FlowOutput) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFlowOutput_FromFlowOutputArnParameters(scope constructs.Construct, id *string, flowOutputArn *string) error {
	return nil
}

func validateFlowOutput_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFlowOutput_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFlowOutput_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFlowOutputParameters(scope constructs.Construct, id *string, props *FlowOutputProps) error {
	return nil
}

