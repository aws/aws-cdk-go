//go:build no_runtime_type_checking

package mixinsawsappstream

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAppBlockPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAppBlockPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAppBlockPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAppBlockPropsMixinParameters(props *CfnAppBlockMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

