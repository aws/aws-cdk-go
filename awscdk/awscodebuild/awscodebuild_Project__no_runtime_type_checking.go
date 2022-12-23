//go:build no_runtime_type_checking

package awscodebuild

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Project) validateAddFileSystemLocationParameters(fileSystemLocation IFileSystemLocation) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddSecondaryArtifactParameters(secondaryArtifact IArtifacts) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddSecondarySourceParameters(secondarySource ISource) error {
	return nil
}

func (p *jsiiProxy_Project) validateAddToRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_Project) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_Project) validateBindAsNotificationRuleSourceParameters(_scope constructs.Construct) error {
	return nil
}

func (p *jsiiProxy_Project) validateBindToCodePipelineParameters(_scope constructs.Construct, options *BindToCodePipelineOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_Project) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_Project) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateMetricBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateMetricDurationParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateMetricFailedBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateMetricSucceededBuildsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateNotifyOnParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *ProjectNotifyOnOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateNotifyOnBuildFailedParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateNotifyOnBuildSucceededParameters(id *string, target awscodestarnotifications.INotificationRuleTarget, options *awscodestarnotifications.NotificationRuleOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnBuildFailedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnBuildStartedParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnBuildSucceededParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnPhaseChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (p *jsiiProxy_Project) validateOnStateChangeParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func validateProject_FromProjectArnParameters(scope constructs.Construct, id *string, projectArn *string) error {
	return nil
}

func validateProject_FromProjectNameParameters(scope constructs.Construct, id *string, projectName *string) error {
	return nil
}

func validateProject_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProject_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateProject_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateProject_SerializeEnvVariablesParameters(environmentVariables *map[string]*BuildEnvironmentVariable) error {
	return nil
}

func validateNewProjectParameters(scope constructs.Construct, id *string, props *ProjectProps) error {
	return nil
}

