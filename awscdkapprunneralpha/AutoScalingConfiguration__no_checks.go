//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoScalingConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AutoScalingConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AutoScalingConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAutoScalingConfiguration_FromArnParameters(scope constructs.Construct, id *string, autoScalingConfigurationArn *string) error {
	return nil
}

func validateAutoScalingConfiguration_FromAutoScalingConfigurationAttributesParameters(scope constructs.Construct, id *string, attrs *AutoScalingConfigurationAttributes) error {
	return nil
}

func validateAutoScalingConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAutoScalingConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAutoScalingConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAutoScalingConfigurationParameters(scope constructs.Construct, id *string, props *AutoScalingConfigurationProps) error {
	return nil
}

