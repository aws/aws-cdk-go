//go:build no_runtime_type_checking

package awsefs

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IFileSystem) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IFileSystem) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

