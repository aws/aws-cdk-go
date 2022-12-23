//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsOptimizedImage) validateGetImageParameters(scope constructs.Construct) error {
	return nil
}

func validateEcsOptimizedImage_AmazonLinuxParameters(options *EcsOptimizedImageOptions) error {
	return nil
}

func validateEcsOptimizedImage_AmazonLinux2Parameters(options *EcsOptimizedImageOptions) error {
	return nil
}

func validateEcsOptimizedImage_WindowsParameters(windowsVersion WindowsOptimizedVersion, options *EcsOptimizedImageOptions) error {
	return nil
}

