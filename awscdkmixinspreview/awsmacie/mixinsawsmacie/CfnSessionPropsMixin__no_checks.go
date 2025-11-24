//go:build no_runtime_type_checking

package mixinsawsmacie

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSessionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSessionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSessionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSessionPropsMixinParameters(props *CfnSessionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

