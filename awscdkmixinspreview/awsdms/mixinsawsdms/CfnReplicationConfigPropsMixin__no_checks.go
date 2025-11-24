//go:build no_runtime_type_checking

package mixinsawsdms

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicationConfigPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReplicationConfigPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReplicationConfigPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReplicationConfigPropsMixinParameters(props *CfnReplicationConfigMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

