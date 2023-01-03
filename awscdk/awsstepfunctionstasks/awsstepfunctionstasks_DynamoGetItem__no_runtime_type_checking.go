//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoGetItem) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateAddPrefixParameters(x *string) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (d *jsiiProxy_DynamoGetItem) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateDynamoGetItem_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateDynamoGetItem_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoGetItem_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateDynamoGetItem_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDynamoGetItem_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewDynamoGetItemParameters(scope constructs.Construct, id *string, props *DynamoGetItemProps) error {
	return nil
}

