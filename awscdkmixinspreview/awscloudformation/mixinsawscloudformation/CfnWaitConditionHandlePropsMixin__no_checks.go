//go:build no_runtime_type_checking

package mixinsawscloudformation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWaitConditionHandlePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWaitConditionHandlePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWaitConditionHandlePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWaitConditionHandlePropsMixinParameters(props *CfnWaitConditionHandleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

