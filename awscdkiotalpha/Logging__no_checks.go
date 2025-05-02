//go:build no_runtime_type_checking

package awscdkiotalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Logging) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_Logging) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_Logging) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLogging_FromLogIdParameters(scope constructs.Construct, id *string, logId *string) error {
	return nil
}

func validateLogging_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLogging_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLogging_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLoggingParameters(scope constructs.Construct, id *string, props *LoggingProps) error {
	return nil
}

