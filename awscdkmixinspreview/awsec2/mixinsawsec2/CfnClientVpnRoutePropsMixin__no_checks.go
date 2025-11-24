//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnClientVpnRoutePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnClientVpnRoutePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnClientVpnRoutePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnClientVpnRoutePropsMixinParameters(props *CfnClientVpnRouteMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

