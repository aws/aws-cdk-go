//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorEcrImageScanAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorEcrImageScanAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorEcrImageScanAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (i *jsiiProxy_InspectorEcrImageScanAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewInspectorEcrImageScanActionParameters(props *InspectorEcrImageScanActionProps) error {
	return nil
}

