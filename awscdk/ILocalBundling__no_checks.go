//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILocalBundling) validateTryBundleParameters(outputDir *string, options *BundlingOptions) error {
	return nil
}

