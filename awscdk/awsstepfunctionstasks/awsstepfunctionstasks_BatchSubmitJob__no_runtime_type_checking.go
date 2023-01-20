//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchSubmitJob) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateAddPrefixParameters(x *string) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (b *jsiiProxy_BatchSubmitJob) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateBatchSubmitJob_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateBatchSubmitJob_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateBatchSubmitJob_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateBatchSubmitJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBatchSubmitJob_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewBatchSubmitJobParameters(scope constructs.Construct, id *string, props *BatchSubmitJobProps) error {
	return nil
}

