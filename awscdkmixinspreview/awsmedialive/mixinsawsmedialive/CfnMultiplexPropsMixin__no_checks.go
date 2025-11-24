//go:build no_runtime_type_checking

package mixinsawsmedialive

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMultiplexPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMultiplexPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMultiplexPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMultiplexPropsMixinParameters(props *CfnMultiplexMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

