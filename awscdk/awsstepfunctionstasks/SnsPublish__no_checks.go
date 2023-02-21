//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsPublish) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateSnsPublish_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateSnsPublish_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateSnsPublish_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateSnsPublish_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSnsPublish_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewSnsPublishParameters(scope constructs.Construct, id *string, props *SnsPublishProps) error {
	return nil
}

