//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsSendMessage) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateAddPrefixParameters(x *string) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateSqsSendMessage_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateSqsSendMessage_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateSqsSendMessage_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateSqsSendMessage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSqsSendMessage_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewSqsSendMessageParameters(scope constructs.Construct, id *string, props *SqsSendMessageProps) error {
	return nil
}

