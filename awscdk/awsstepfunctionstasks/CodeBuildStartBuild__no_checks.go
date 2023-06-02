//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildStartBuild) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateAddPrefixParameters(x *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateCodeBuildStartBuild_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateCodeBuildStartBuild_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateCodeBuildStartBuild_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateCodeBuildStartBuild_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCodeBuildStartBuild_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewCodeBuildStartBuildParameters(scope constructs.Construct, id *string, props *CodeBuildStartBuildProps) error {
	return nil
}

