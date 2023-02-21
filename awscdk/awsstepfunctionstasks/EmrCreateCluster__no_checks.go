//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrCreateCluster) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateAddPrefixParameters(x *string) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (e *jsiiProxy_EmrCreateCluster) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateEmrCreateCluster_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateEmrCreateCluster_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEmrCreateCluster_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateEmrCreateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEmrCreateCluster_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewEmrCreateClusterParameters(scope constructs.Construct, id *string, props *EmrCreateClusterProps) error {
	return nil
}

