//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Workflow) validateAddConditionalTriggerParameters(id *string, options *ConditionalTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateAddCustomScheduledTriggerParameters(id *string, options *CustomScheduledTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateAddDailyScheduledTriggerParameters(id *string, options *DailyScheduleTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateAddNotifyEventTriggerParameters(id *string, options *NotifyEventTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateAddOnDemandTriggerParameters(id *string, options *OnDemandTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateAddWeeklyScheduledTriggerParameters(id *string, options *WeeklyScheduleTriggerOptions) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateBuildWorkflowArnParameters(scope constructs.Construct, workflowName *string) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (w *jsiiProxy_Workflow) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateWorkflow_ExtractNameFromArnParameters(scope constructs.Construct, workflowArn *string) error {
	return nil
}

func validateWorkflow_FromWorkflowArnParameters(scope constructs.Construct, id *string, workflowArn *string) error {
	return nil
}

func validateWorkflow_FromWorkflowAttributesParameters(scope constructs.Construct, id *string, attrs *WorkflowAttributes) error {
	return nil
}

func validateWorkflow_FromWorkflowNameParameters(scope constructs.Construct, id *string, workflowName *string) error {
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

func validateNewWorkflowParameters(scope constructs.Construct, id *string, props *WorkflowProps) error {
	return nil
}

