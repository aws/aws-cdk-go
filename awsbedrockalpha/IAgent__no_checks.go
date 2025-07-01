//go:build no_runtime_type_checking

package awsbedrockalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAgent) validateGrantInvokeParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAgent) validateMetricCountParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (i *jsiiProxy_IAgent) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IAgent) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

