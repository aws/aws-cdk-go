//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorSourceCodeScanAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorSourceCodeScanAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorSourceCodeScanAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (i *jsiiProxy_InspectorSourceCodeScanAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewInspectorSourceCodeScanActionParameters(props *InspectorSourceCodeScanActionProps) error {
	return nil
}

