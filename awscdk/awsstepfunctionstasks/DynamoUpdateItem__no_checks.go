//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoUpdateItem) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateAddPrefixParameters(x *string) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DynamoUpdateItem) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateDynamoUpdateItem_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateDynamoUpdateItem_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoUpdateItem_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoUpdateItem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDynamoUpdateItem_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewDynamoUpdateItemParameters(scope constructs.Construct, id *string, props *DynamoUpdateItemProps) error {
	return nil
}

