//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SparkJob) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateBuildJobArnParameters(scope constructs.Construct, jobName *string) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateCodeS3ObjectUrlParameters(code Code) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateMetricParameters(metricName *string, type_ MetricType, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateMetricFailureParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateMetricSuccessParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateMetricTimeoutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateNonExecutableCommonArgumentsParameters(props *SparkJobProps) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateOnFailureParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateOnStateChangeParameters(id *string, jobState JobState, options *awsevents.OnEventOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateOnSuccessParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateOnTimeoutParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateSetupContinuousLoggingParameters(role awsiam.IRole, props *ContinuousLoggingProps) error {
	return nil
}

func (s *jsiiProxy_SparkJob) validateSetupExtraCodeArgumentsParameters(args *map[string]*string, props *SparkExtraCodeProps) error {
	return nil
}

func validateSparkJob_FromJobAttributesParameters(scope constructs.Construct, id *string, attrs *JobAttributes) error {
	return nil
}

func validateSparkJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSparkJob_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSparkJob_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSparkJobParameters(scope constructs.Construct, id *string, props *SparkJobProps) error {
	return nil
}

