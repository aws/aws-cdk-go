//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSecurityGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSecurityGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSecurityGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSecurityGroupPropsMixinParameters(props *CfnSecurityGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

