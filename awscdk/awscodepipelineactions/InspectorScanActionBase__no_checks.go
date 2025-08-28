//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InspectorScanActionBase) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorScanActionBase) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (i *jsiiProxy_InspectorScanActionBase) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (i *jsiiProxy_InspectorScanActionBase) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewInspectorScanActionBaseParameters(props *InspectorScanActionBaseProps) error {
	return nil
}

