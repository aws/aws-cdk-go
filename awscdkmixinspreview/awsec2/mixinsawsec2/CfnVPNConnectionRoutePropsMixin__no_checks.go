//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPNConnectionRoutePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPNConnectionRoutePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPNConnectionRoutePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPNConnectionRoutePropsMixinParameters(props *CfnVPNConnectionRouteMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

