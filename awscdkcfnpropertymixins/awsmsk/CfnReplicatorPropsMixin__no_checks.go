//go:build no_runtime_type_checking

package awsmsk

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicatorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReplicatorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReplicatorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReplicatorPropsMixinParameters(props *CfnReplicatorMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

