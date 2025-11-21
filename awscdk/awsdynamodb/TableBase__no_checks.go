//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantOnKeyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantStreamParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantStreamReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantTableListStreamsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricConditionalCheckFailedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricConsumedReadCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricConsumedWriteCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricSuccessfulRequestLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricSystemErrorsForOperationsParameters(props *SystemErrorsForOperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricThrottledRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricThrottledRequestsForOperationParameters(operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricThrottledRequestsForOperationsParameters(props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTableBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTableBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTableBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

