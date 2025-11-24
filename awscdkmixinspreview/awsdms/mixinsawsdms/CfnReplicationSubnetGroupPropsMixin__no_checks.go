//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicationSubnetGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReplicationSubnetGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReplicationSubnetGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReplicationSubnetGroupPropsMixinParameters(props *CfnReplicationSubnetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

