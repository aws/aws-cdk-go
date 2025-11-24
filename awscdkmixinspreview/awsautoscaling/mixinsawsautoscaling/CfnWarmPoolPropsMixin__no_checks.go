//go:build no_runtime_type_checking

package mixinsawsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWarmPoolPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWarmPoolPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWarmPoolPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWarmPoolPropsMixinParameters(props *CfnWarmPoolMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

