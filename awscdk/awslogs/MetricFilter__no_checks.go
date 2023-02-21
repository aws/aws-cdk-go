//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetricFilter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (m *jsiiProxy_MetricFilter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (m *jsiiProxy_MetricFilter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (m *jsiiProxy_MetricFilter) validateMetricParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateMetricFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMetricFilter_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateMetricFilter_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewMetricFilterParameters(scope constructs.Construct, id *string, props *MetricFilterProps) error {
	return nil
}

