//go:build no_runtime_type_checking

package awsneptune

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IDatabaseCluster) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IDatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

