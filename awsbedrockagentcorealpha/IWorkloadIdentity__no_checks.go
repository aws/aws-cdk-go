//go:build no_runtime_type_checking

package awsbedrockagentcorealpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IWorkloadIdentity) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IWorkloadIdentity) validateGrantAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IWorkloadIdentity) validateGrantFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IWorkloadIdentity) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IWorkloadIdentity) validateGrantUseParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IWorkloadIdentity) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

