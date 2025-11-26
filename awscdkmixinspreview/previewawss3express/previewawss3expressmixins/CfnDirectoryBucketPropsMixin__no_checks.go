//go:build no_runtime_type_checking

package previewawss3expressmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDirectoryBucketPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDirectoryBucketPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDirectoryBucketPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDirectoryBucketPropsMixinParameters(props *CfnDirectoryBucketMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

