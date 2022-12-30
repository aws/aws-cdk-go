//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Task) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddCatchParameters(handler IChainable, props *CatchProps) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddPrefixParameters(x *string) error {
	return nil
}

func (t *jsiiProxy_Task) validateAddRetryParameters(props *RetryProps) error {
	return nil
}

func (t *jsiiProxy_Task) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (t *jsiiProxy_Task) validateMakeDefaultParameters(def State) error {
	return nil
}

func (t *jsiiProxy_Task) validateMakeNextParameters(next State) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Task) validateNextParameters(next IChainable) error {
	return nil
}

func (t *jsiiProxy_Task) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_Task) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_Task) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateTask_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateTask_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateTask_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateTask_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTask_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewTaskParameters(scope constructs.Construct, id *string, props *TaskProps) error {
	return nil
}

