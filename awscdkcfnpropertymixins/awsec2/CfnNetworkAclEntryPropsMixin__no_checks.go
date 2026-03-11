//go:build no_runtime_type_checking

package awsec2

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

func validateNewCfnNetworkAclEntryPropsMixinParameters(props *CfnNetworkAclEntryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

