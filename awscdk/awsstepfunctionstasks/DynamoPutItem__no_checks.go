//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoPutItem) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddItemProcessorParameters(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddPrefixParameters(x *string) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateNextParameters(state awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DynamoPutItem) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateDynamoPutItem_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateDynamoPutItem_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoPutItem_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoPutItem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDynamoPutItem_JsonataParameters(scope constructs.Construct, id *string, props *DynamoPutItemJsonataProps) error {
	return nil
}

func validateDynamoPutItem_JsonPathParameters(scope constructs.Construct, id *string, props *DynamoPutItemJsonPathProps) error {
	return nil
}

func validateDynamoPutItem_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func (j *jsiiProxy_DynamoPutItem) validateSetProcessorConfigParameters(val *awsstepfunctions.ProcessorConfig) error {
	return nil
}

func validateNewDynamoPutItemParameters(scope constructs.Construct, id *string, props *DynamoPutItemProps) error {
	return nil
}

