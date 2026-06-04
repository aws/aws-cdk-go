//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvaluatorBase) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (e *jsiiProxy_EvaluatorBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_EvaluatorBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_EvaluatorBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_EvaluatorBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEvaluatorBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEvaluatorBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEvaluatorBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEvaluatorBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

