//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FlowLog) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (f *jsiiProxy_FlowLog) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (f *jsiiProxy_FlowLog) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateFlowLog_FromFlowLogIdParameters(scope constructs.Construct, id *string, flowLogId *string) error {
	return nil
}

func validateFlowLog_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFlowLog_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateFlowLog_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewFlowLogParameters(scope constructs.Construct, id *string, props *FlowLogProps) error {
	return nil
}

