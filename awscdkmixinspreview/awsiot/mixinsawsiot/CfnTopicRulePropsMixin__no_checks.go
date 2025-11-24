//go:build no_runtime_type_checking

package mixinsawsiot

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTopicRulePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTopicRulePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTopicRulePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTopicRulePropsMixinParameters(props *CfnTopicRuleMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

