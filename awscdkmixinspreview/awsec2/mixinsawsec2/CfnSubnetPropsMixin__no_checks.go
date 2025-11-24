//go:build no_runtime_type_checking

package mixinsawsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnSubnetPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnSubnetPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnSubnetPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnSubnetPropsMixinParameters(props *CfnSubnetMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

