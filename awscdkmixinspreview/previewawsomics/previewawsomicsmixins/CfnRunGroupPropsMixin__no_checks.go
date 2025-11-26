//go:build no_runtime_type_checking

package previewawsomicsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRunGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRunGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRunGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRunGroupPropsMixinParameters(props *CfnRunGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

