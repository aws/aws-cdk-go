//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsRunTask) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateAddPrefixParameters(x *string) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateEcsRunTask_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateEcsRunTask_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEcsRunTask_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEcsRunTask_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEcsRunTask_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewEcsRunTaskParameters(scope constructs.Construct, id *string, props *EcsRunTaskProps) error {
	return nil
}

