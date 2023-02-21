//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoDeleteItem) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateAddPrefixParameters(x *string) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DynamoDeleteItem) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateDynamoDeleteItem_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateDynamoDeleteItem_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoDeleteItem_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoDeleteItem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDynamoDeleteItem_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewDynamoDeleteItemParameters(scope constructs.Construct, id *string, props *DynamoDeleteItemProps) error {
	return nil
}

