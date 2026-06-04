//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Evaluator) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (e *jsiiProxy_Evaluator) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Evaluator) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Evaluator) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Evaluator) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEvaluator_FromEvaluatorArnParameters(scope constructs.Construct, id *string, evaluatorArn *string) error {
	return nil
}

func validateEvaluator_FromEvaluatorAttributesParameters(scope constructs.Construct, id *string, attrs *EvaluatorAttributes) error {
	return nil
}

func validateEvaluator_FromEvaluatorIdParameters(scope constructs.Construct, id *string, evaluatorId *string) error {
	return nil
}

func validateEvaluator_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEvaluator_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEvaluator_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewEvaluatorParameters(scope constructs.Construct, id *string, props *EvaluatorProps) error {
	return nil
}

