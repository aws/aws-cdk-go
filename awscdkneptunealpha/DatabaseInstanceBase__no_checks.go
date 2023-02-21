//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstanceBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstanceBase) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateDatabaseInstanceBase_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstanceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstanceBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstanceBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

