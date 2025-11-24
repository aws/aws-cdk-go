//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnNetworkAclEntryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnNetworkAclEntryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnNetworkAclEntryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnNetworkAclEntryPropsMixinParameters(props *CfnNetworkAclEntryMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

