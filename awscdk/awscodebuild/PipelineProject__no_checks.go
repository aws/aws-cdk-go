//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineProject) validateAddFileSystemLocationParameters(fileSystemLocation IFileSystemLocation) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateAddSecondaryArtifactParameters(secondaryArtifact IArtifacts) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateAddSecondarySourceParameters(secondarySource ISource) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateBindAsNotificationRuleSourceParameters(_scope constructs.Construct) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateBindToCodePipelineParameters(_scope constructs.Construct, options *BindToCodePipelineOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateMetricBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateMetricFailedBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateMetricSucceededBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateNotifyOnBuildFailedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateNotifyOnBuildSucceededParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnBuildFailedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnBuildStartedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnBuildSucceededParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnPhaseChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineProject) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validatePipelineProject_FromProjectArnParameters(scope constructs.Construct, id *string, projectArn *string) error {
	return nil
}

func validatePipelineProject_FromProjectNameParameters(scope constructs.Construct, id *string, projectName *string) error {
	return nil
}

func validatePipelineProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePipelineProject_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePipelineProject_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePipelineProject_SerializeEnvVariablesParameters(environmentVariables *map[string]*BuildEnvironmentVariable) error {
	return nil
}

func validateNewPipelineProjectParameters(scope constructs.Construct, id *string, props *PipelineProjectProps) error {
	return nil
}

