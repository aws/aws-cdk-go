//go:build no_runtime_type_checking

package mixinsawsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInputPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInputPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInputPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInputPropsMixinParameters(props *CfnInputMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

