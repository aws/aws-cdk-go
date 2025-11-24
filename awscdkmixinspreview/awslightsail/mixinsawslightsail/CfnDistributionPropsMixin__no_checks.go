//go:build no_runtime_type_checking

package mixinsawslightsail

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDistributionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDistributionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDistributionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDistributionPropsMixinParameters(props *CfnDistributionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

