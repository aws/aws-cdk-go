//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvoke) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateAddPrefixParameters(x *string) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateLambdaInvoke_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateLambdaInvoke_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateLambdaInvoke_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateLambdaInvoke_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLambdaInvoke_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewLambdaInvokeParameters(scope constructs.Construct, id *string, props *LambdaInvokeProps) error {
	return nil
}

