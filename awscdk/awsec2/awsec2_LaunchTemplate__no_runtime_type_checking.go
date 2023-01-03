//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplate) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplate) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplate) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLaunchTemplate_FromLaunchTemplateAttributesParameters(scope constructs.Construct, id *string, attrs *LaunchTemplateAttributes) error {
	return nil
}

func validateLaunchTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLaunchTemplate_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLaunchTemplate_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLaunchTemplateParameters(scope constructs.Construct, id *string, props *LaunchTemplateProps) error {
	return nil
}

