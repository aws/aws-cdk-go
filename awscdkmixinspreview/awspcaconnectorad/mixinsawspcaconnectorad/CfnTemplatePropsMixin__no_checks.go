//go:build no_runtime_type_checking

package mixinsawspcaconnectorad

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTemplatePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTemplatePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTemplatePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTemplatePropsMixinParameters(props *CfnTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

