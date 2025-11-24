//go:build no_runtime_type_checking

package mixinsawsvpclattice

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnServiceNetworkPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnServiceNetworkPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnServiceNetworkPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnServiceNetworkPropsMixinParameters(props *CfnServiceNetworkMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

