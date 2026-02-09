//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IGatewayTarget) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGatewayTarget) validateGrantManageParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGatewayTarget) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IGatewayTarget) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

