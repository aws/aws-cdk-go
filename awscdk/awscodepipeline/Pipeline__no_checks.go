//go:build no_runtime_type_checking

package awscodepipeline

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pipeline) validateAddStageParameters(props *StageOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateBindAsNotificationRuleSourceParameters(_scope constructs.Construct) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *PipelineNotifyOnOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateNotifyOnAnyActionStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateNotifyOnAnyManualApprovalStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateNotifyOnAnyStageStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateNotifyOnExecutionStateChangeParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Pipeline) validateStageParameters(stageName *string) error {
	return nil
}

func validatePipeline_FromPipelineArnParameters(scope constructs.Construct, id *string, pipelineArn *string) error {
	return nil
}

func validatePipeline_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipeline_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePipeline_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPipelineParameters(scope constructs.Construct, id *string, props *PipelineProps) error {
	return nil
}

