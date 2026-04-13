//go:build no_runtime_type_checking

package awssecurityagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTargetDomainPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTargetDomainPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTargetDomainPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTargetDomainPropsMixinParameters(props *CfnTargetDomainMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

