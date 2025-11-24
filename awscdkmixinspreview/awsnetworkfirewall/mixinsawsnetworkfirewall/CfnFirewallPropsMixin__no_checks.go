//go:build no_runtime_type_checking

package mixinsawsnetworkfirewall

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnFirewallPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnFirewallPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnFirewallPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnFirewallPropsMixinParameters(props *CfnFirewallMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

