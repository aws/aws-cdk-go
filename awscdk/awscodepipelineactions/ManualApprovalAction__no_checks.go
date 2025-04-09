//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManualApprovalAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalAction) validateGrantManualApprovalParameters(grantable awsiam.IGrantable) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewManualApprovalActionParameters(props *ManualApprovalActionProps) error {
	return nil
}

