//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrAddStep) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateAddPrefixParameters(x *string) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (e *jsiiProxy_EmrAddStep) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateEmrAddStep_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateEmrAddStep_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEmrAddStep_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEmrAddStep_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEmrAddStep_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewEmrAddStepParameters(scope constructs.Construct, id *string, props *EmrAddStepProps) error {
	return nil
}

