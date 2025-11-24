//go:build no_runtime_type_checking

package mixinsawsrds

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDBSecurityGroupPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDBSecurityGroupPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDBSecurityGroupPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDBSecurityGroupPropsMixinParameters(props *CfnDBSecurityGroupMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

