//go:build no_runtime_type_checking

package awsinspector

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AssessmentTemplate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AssessmentTemplate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AssessmentTemplate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAssessmentTemplate_FromCfnAssessmentTemplateParameters(scope constructs.Construct, id *string, template CfnAssessmentTemplate) error {
	return nil
}

func validateAssessmentTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAssessmentTemplate_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAssessmentTemplate_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAssessmentTemplateParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

