//go:build no_runtime_type_checking

package awss3tables

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTableBucketPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTableBucketPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTableBucketPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTableBucketPropsMixinParameters(props *CfnTableBucketMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

