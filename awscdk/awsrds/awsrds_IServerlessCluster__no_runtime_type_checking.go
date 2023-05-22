//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IServerlessCluster) validateGrantDataApiAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IServerlessCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

