//go:build no_runtime_type_checking

package mixinsawsmemorydb

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubnetGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubnetGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubnetGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubnetGroupPropsMixinParameters(props *CfnSubnetGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

