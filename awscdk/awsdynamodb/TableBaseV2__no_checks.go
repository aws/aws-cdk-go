//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableBaseV2) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantOnKeyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantStreamParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantStreamReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantTableListStreamsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateGrantWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricConditionalCheckFailedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricConsumedReadCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricConsumedWriteCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricSuccessfulRequestLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricSystemErrorsForOperationsParameters(props *SystemErrorsForOperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricThrottledRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricThrottledRequestsForOperationParameters(operation *string, props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricThrottledRequestsForOperationsParameters(props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableBaseV2) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTableBaseV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTableBaseV2_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTableBaseV2_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableBaseV2Parameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

