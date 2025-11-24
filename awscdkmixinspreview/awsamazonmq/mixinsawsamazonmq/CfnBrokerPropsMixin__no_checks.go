//go:build no_runtime_type_checking

package mixinsawsamazonmq

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBrokerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBrokerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBrokerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBrokerPropsMixinParameters(props *CfnBrokerMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

