//go:build no_runtime_type_checking

package mixinsawsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnQueueEnvironmentPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnQueueEnvironmentPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnQueueEnvironmentPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnQueueEnvironmentPropsMixinParameters(props *CfnQueueEnvironmentMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

