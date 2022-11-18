//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Connection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Connection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateConnection_FromConnectionAttributesParameters(scope constructs.Construct, id *string, attrs *ConnectionAttributes) error {
	return nil
}

func validateConnection_FromEventBusArnParameters(scope constructs.Construct, id *string, connectionArn *string, connectionSecretArn *string) error {
	return nil
}

func validateConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConnection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConnection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConnectionParameters(scope constructs.Construct, id *string, props *ConnectionProps) error {
	return nil
}

