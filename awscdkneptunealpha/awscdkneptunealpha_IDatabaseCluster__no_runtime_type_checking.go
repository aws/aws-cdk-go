//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Neptune
package awscdkneptunealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseCluster) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

