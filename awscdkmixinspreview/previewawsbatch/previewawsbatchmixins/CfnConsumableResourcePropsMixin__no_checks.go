//go:build no_runtime_type_checking

package previewawsbatchmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConsumableResourcePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConsumableResourcePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConsumableResourcePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConsumableResourcePropsMixinParameters(props *CfnConsumableResourceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

