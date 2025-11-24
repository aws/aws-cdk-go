//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPNConnectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPNConnectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPNConnectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPNConnectionPropsMixinParameters(props *CfnVPNConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

