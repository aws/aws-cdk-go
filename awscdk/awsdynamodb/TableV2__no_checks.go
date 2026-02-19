//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableV2) validateAddGlobalSecondaryIndexParameters(props *GlobalSecondaryIndexPropsV2) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateAddLocalSecondaryIndexParameters(props *LocalSecondaryIndexProps) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateAddReplicaParameters(props *ReplicaTableProps) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantOnKeyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantStreamParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantStreamReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantTableListStreamsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateGrantWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricConditionalCheckFailedRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricConsumedReadCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricConsumedWriteCapacityUnitsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricSuccessfulRequestLatencyParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricSystemErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricSystemErrorsForOperationsParameters(props *SystemErrorsForOperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricThrottledRequestsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricThrottledRequestsForOperationParameters(operation *string, props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricThrottledRequestsForOperationsParameters(props *OperationsMetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateMetricUserErrorsParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (t *jsiiProxy_TableV2) validateReplicaParameters(region *string) error {
	return nil
}

func validateTableV2_FromTableArnParameters(scope constructs.Construct, id *string, tableArn *string) error {
	return nil
}

func validateTableV2_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributesV2) error {
	return nil
}

func validateTableV2_FromTableNameParameters(scope constructs.Construct, id *string, tableName *string) error {
	return nil
}

func validateTableV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTableV2_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTableV2_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableV2Parameters(scope constructs.Construct, id *string, props *TablePropsV2) error {
	return nil
}

