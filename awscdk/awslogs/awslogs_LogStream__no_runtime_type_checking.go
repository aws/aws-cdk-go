//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogStream) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (l *jsiiProxy_LogStream) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (l *jsiiProxy_LogStream) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateLogStream_FromLogStreamNameParameters(scope constructs.Construct, id *string, logStreamName *string) error {
	return nil
}

func validateLogStream_IsConstructParameters(x interface{}) error {
	return nil
}

func validateLogStream_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateLogStream_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewLogStreamParameters(scope constructs.Construct, id *string, props *LogStreamProps) error {
	return nil
}

