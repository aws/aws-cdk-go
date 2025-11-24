//go:build no_runtime_type_checking

package mixinsawsconfig

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnConformancePackPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnConformancePackPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnConformancePackPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnConformancePackPropsMixinParameters(props *CfnConformancePackMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

