//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Table) validateAddGlobalSecondaryIndexParameters(props *GlobalSecondaryIndexProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateAddLocalSecondaryIndexParameters(props *LocalSecondaryIndexProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Table) validateAutoScaleGlobalSecondaryIndexReadCapacityParameters(indexName *string, props *EnableScalingProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateAutoScaleGlobalSecondaryIndexWriteCapacityParameters(indexName *string, props *EnableScalingProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateAutoScaleReadCapacityParameters(props *EnableScalingProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateAutoScaleWriteCapacityParameters(props *EnableScalingProps) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantStreamParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantStreamReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantTableListStreamsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricConditionalCheckFailedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricConsumedReadCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricConsumedWriteCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricSuccessfulRequestLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricSystemErrorsForOperationsParameters(props *SystemErrorsForOperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricThrottledRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricThrottledRequestsForOperationParameters(operation *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricThrottledRequestsForOperationsParameters(props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_Table) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateTable_FromTableArnParameters(scope constructs.Construct, id *string, tableArn *string) error {
	return nil
}

func validateTable_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributes) error {
	return nil
}

func validateTable_FromTableNameParameters(scope constructs.Construct, id *string, tableName *string) error {
	return nil
}

func validateTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTable_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTable_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableParameters(scope constructs.Construct, id *string, props *TableProps) error {
	return nil
}

