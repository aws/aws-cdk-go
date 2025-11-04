//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockInvokeModel) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddItemProcessorParameters(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddPrefixParameters(x *string) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateNextParameters(state awsstepfunctions.IChainable) error {
	return nil
}

func (b *jsiiProxy_BedrockInvokeModel) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateBedrockInvokeModel_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateBedrockInvokeModel_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateBedrockInvokeModel_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateBedrockInvokeModel_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBedrockInvokeModel_JsonataParameters(scope constructs.Construct, id *string, props *BedrockInvokeModelJsonataProps) error {
	return nil
}

func validateBedrockInvokeModel_JsonPathParameters(scope constructs.Construct, id *string, props *BedrockInvokeModelJsonPathProps) error {
	return nil
}

func validateBedrockInvokeModel_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_BedrockInvokeModel) validateSetProcessorConfigParameters(val *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func validateNewBedrockInvokeModelParameters(scope constructs.Construct, id *string, props *BedrockInvokeModelProps) error {
	return nil
}

