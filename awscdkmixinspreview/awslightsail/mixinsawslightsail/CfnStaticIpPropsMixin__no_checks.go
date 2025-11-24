//go:build no_runtime_type_checking

package mixinsawslightsail

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStaticIpPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStaticIpPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStaticIpPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStaticIpPropsMixinParameters(props *CfnStaticIpMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

