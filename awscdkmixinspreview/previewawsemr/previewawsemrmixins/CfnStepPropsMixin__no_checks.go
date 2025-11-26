//go:build no_runtime_type_checking

package previewawsemrmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStepPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStepPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStepPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStepPropsMixinParameters(props *CfnStepMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

