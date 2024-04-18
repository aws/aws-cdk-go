//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HttpInvoke) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddItemProcessorParameters(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddPrefixParameters(x *string) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (h *jsiiProxy_HttpInvoke) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateHttpInvoke_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateHttpInvoke_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateHttpInvoke_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateHttpInvoke_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHttpInvoke_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_HttpInvoke) validateSetProcessorConfigParameters(val *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func validateNewHttpInvokeParameters(scope constructs.Construct, id *string, props *HttpInvokeProps) error {
	return nil
}

