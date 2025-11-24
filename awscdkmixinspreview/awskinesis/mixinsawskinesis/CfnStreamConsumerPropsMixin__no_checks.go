//go:build no_runtime_type_checking

package mixinsawskinesis

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStreamConsumerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStreamConsumerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStreamConsumerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStreamConsumerPropsMixinParameters(props *CfnStreamConsumerMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

