//go:build no_runtime_type_checking

package awss3deployment

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketDeployment) validateAddSourceParameters(source ISource) error {
	return nil
}

func validateBucketDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBucketDeploymentParameters(scope constructs.Construct, id *string, props *BucketDeploymentProps) error {
	return nil
}

