//go:build no_runtime_type_checking

package awss3mixins

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketVersioning) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (b *jsiiProxy_BucketVersioning) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucketVersioning_IsMixinParameters(x interface{}) error {
	return nil
}

