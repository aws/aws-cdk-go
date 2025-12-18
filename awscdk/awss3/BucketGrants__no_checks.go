//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketGrants) validateDeleteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validatePutParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validatePutAclParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validateReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validateReadWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validateReplicationPermissionParameters(identity awsiam.IGrantable, props *GrantReplicationPermissionProps) error {
	return nil
}

func (b *jsiiProxy_BucketGrants) validateWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func validateBucketGrants_FromBucketParameters(bucket interfacesawss3.IBucketRef) error {
	return nil
}

