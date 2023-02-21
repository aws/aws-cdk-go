//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBridgePutEvents) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateAddPrefixParameters(x *string) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateEventBridgePutEvents_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateEventBridgePutEvents_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEventBridgePutEvents_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEventBridgePutEvents_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEventBridgePutEvents_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewEventBridgePutEventsParameters(scope constructs.Construct, id *string, props *EventBridgePutEventsProps) error {
	return nil
}

