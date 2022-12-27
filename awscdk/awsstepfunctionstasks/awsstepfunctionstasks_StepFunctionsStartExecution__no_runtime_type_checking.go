//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateStepFunctionsStartExecution_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateStepFunctionsStartExecution_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateStepFunctionsStartExecution_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateStepFunctionsStartExecution_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStepFunctionsStartExecution_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewStepFunctionsStartExecutionParameters(scope constructs.Construct, id *string, props *StepFunctionsStartExecutionProps) error {
	return nil
}

