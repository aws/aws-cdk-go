//go:build no_runtime_type_checking

package awskinesis

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IStreamConsumer) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IStreamConsumer) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStreamConsumer) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IStreamConsumer) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

