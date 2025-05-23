//go:build no_runtime_type_checking

package awskinesis

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StreamConsumer) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_StreamConsumer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_StreamConsumer) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_StreamConsumer) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_StreamConsumer) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_StreamConsumer) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateStreamConsumer_FromStreamConsumerArnParameters(scope constructs.Construct, id *string, streamConsumerArn *string) error {
	return nil
}

func validateStreamConsumer_FromStreamConsumerAttributesParameters(scope constructs.Construct, id *string, attrs *StreamConsumerAttributes) error {
	return nil
}

func validateStreamConsumer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStreamConsumer_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateStreamConsumer_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewStreamConsumerParameters(scope constructs.Construct, id *string, props *StreamConsumerProps) error {
	return nil
}

