//go:build no_runtime_type_checking

package awsstepfunctions

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskStateBase) validateAddBranchParameters(branch StateGraph) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateAddCatchParameters(handler IChainable, props *CatchProps) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateAddChoiceParameters(condition Condition, next State) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateAddIteratorParameters(iteration StateGraph) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateAddPrefixParameters(x *string) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateAddRetryParameters(props *RetryProps) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateBindToGraphParameters(graph StateGraph) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMakeDefaultParameters(def State) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMakeNextParameters(next State) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateNextParameters(next IChainable) error {
	return nil
}

func (t *jsiiProxy_TaskStateBase) validateWhenBoundToGraphParameters(graph StateGraph) error {
	return nil
}

func validateTaskStateBase_FilterNextablesParameters(states *[]State) error {
	return nil
}

func validateTaskStateBase_FindReachableEndStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateTaskStateBase_FindReachableStatesParameters(start State, options *FindStateOptions) error {
	return nil
}

func validateTaskStateBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTaskStateBase_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewTaskStateBaseParameters(scope constructs.Construct, id *string, props *TaskStateBaseProps) error {
	return nil
}

