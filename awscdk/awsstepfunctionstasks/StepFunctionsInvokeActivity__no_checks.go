//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsInvokeActivity) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateStepFunctionsInvokeActivity_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateStepFunctionsInvokeActivity_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateStepFunctionsInvokeActivity_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateStepFunctionsInvokeActivity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStepFunctionsInvokeActivity_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewStepFunctionsInvokeActivityParameters(scope constructs.Construct, id *string, props *StepFunctionsInvokeActivityProps) error {
	return nil
}

