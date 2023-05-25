//go:build no_runtime_type_checking

package awsroute53targets

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketWebsiteTarget) validateBindParameters(_record awsroute53.IRecordSet) error {
	return nil
}

func validateNewBucketWebsiteTargetParameters(bucket awss3.IBucket) error {
	return nil
}

