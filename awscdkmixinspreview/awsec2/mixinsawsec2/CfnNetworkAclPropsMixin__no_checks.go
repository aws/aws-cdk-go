//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNetworkAclPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNetworkAclPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNetworkAclPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNetworkAclPropsMixinParameters(props *CfnNetworkAclMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

