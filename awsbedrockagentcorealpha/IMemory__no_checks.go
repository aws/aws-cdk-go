//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IMemory) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantDeleteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantDeleteLongTermMemoryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantDeleteShortTermMemoryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantReadLongTermMemoryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantReadShortTermMemoryParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricParameters(metricName *string, dimensions *map[string]*string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricErrorsForApiOperationParameters(operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricEventCreationCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricForApiOperationParameters(metricName *string, operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricInvocationsForApiOperationParameters(operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateMetricLatencyForApiOperationParameters(operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IMemory) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

