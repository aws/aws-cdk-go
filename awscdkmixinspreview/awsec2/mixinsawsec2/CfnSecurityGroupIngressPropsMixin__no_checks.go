//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSecurityGroupIngressPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSecurityGroupIngressPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSecurityGroupIngressPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSecurityGroupIngressPropsMixinParameters(props *CfnSecurityGroupIngressMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

