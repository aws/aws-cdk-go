//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EvaluateExpression) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateAddPrefixParameters(x *string) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (e *jsiiProxy_EvaluateExpression) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateEvaluateExpression_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateEvaluateExpression_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEvaluateExpression_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEvaluateExpression_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEvaluateExpression_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewEvaluateExpressionParameters(scope constructs.Construct, id *string, props *EvaluateExpressionProps) error {
	return nil
}

