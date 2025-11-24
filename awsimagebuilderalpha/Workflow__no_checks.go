//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Workflow) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateWorkflow_FromWorkflowArnParameters(scope constructs.Construct, id *string, workflowArn *string) error {
	return nil
}

func validateWorkflow_FromWorkflowAttributesParameters(scope constructs.Construct, id *string, attrs *WorkflowAttributes) error {
	return nil
}

func validateWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkflow_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkflow_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateWorkflow_IsWorkflowParameters(x interface{}) error {
	return nil
}

func validateNewWorkflowParameters(scope constructs.Construct, id *string, props *WorkflowProps) error {
	return nil
}

