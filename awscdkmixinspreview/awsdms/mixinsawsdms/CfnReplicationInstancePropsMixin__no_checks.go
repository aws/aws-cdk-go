//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicationInstancePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReplicationInstancePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReplicationInstancePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReplicationInstancePropsMixinParameters(props *CfnReplicationInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

