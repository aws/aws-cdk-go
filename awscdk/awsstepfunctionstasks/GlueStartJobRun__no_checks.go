//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueStartJobRun) validateAddBranchParameters(branch awsstepfunctions.StateGraph) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateAddCatchParameters(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateAddChoiceParameters(condition awsstepfunctions.Condition, next awsstepfunctions.State) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateAddIteratorParameters(iteration awsstepfunctions.StateGraph) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateAddPrefixParameters(x *string) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateAddRetryParameters(props *awsstepfunctions.RetryProps) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateBindToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMakeDefaultParameters(def awsstepfunctions.State) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMakeNextParameters(next awsstepfunctions.State) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricFailedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricHeartbeatTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricRunTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricScheduledParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricScheduleTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricStartedParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricSucceededParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricTimeParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateMetricTimedOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateNextParameters(next awsstepfunctions.IChainable) error {
	return nil
}

func (g *jsiiProxy_GlueStartJobRun) validateWhenBoundToGraphParameters(graph awsstepfunctions.StateGraph) error {
	return nil
}

func validateGlueStartJobRun_FilterNextablesParameters(states *[]awsstepfunctions.State) error {
	return nil
}

func validateGlueStartJobRun_FindReachableEndStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateGlueStartJobRun_FindReachableStatesParameters(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) error {
	return nil
}

func validateGlueStartJobRun_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlueStartJobRun_PrefixStatesParameters(root constructs.IConstruct, prefix *string) error {
	return nil
}

func validateNewGlueStartJobRunParameters(scope constructs.Construct, id *string, props *GlueStartJobRunProps) error {
	return nil
}

