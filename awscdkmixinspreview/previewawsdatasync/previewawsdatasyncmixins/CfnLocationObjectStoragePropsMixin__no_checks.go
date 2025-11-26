//go:build no_runtime_type_checking

package previewawsdatasyncmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLocationObjectStoragePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLocationObjectStoragePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLocationObjectStoragePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLocationObjectStoragePropsMixinParameters(props *CfnLocationObjectStorageMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

