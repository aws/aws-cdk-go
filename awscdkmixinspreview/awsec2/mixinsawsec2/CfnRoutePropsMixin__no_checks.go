//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRoutePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRoutePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRoutePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRoutePropsMixinParameters(props *CfnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

