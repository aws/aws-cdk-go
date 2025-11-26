//go:build no_runtime_type_checking

package previewawsdetectivemixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGraphPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGraphPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGraphPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGraphPropsMixinParameters(props *CfnGraphMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

