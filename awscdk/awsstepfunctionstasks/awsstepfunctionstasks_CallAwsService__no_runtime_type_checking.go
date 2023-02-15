//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CallAwsService) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateAddPrefixParameters(x *string) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (c *jsiiProxy_CallAwsService) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateCallAwsService_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateCallAwsService_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateCallAwsService_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateCallAwsService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCallAwsService_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewCallAwsServiceParameters(scope constructs.Construct, id *string, props *CallAwsServiceProps) error {
	return nil
}

