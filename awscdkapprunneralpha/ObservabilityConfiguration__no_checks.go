//go:build no_runtime_type_checking

package awscdkapprunneralpha

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_ObservabilityConfiguration) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_ObservabilityConfiguration) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_ObservabilityConfiguration) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateObservabilityConfiguration_FromArnParameters(scope constructs.Construct, id *string, observabilityConfigurationArn *string) error {
	return nil
}

func validateObservabilityConfiguration_FromObservabilityConfigurationAttributesParameters(scope constructs.Construct, id *string, attrs *ObservabilityConfigurationAttributes) error {
	return nil
}

func validateObservabilityConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateObservabilityConfiguration_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateObservabilityConfiguration_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewObservabilityConfigurationParameters(scope constructs.Construct, id *string, props *ObservabilityConfigurationProps) error {
	return nil
}

