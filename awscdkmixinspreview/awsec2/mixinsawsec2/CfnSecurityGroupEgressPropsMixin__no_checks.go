//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSecurityGroupEgressPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSecurityGroupEgressPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSecurityGroupEgressPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSecurityGroupEgressPropsMixinParameters(props *CfnSecurityGroupEgressMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

