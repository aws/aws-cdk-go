//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnParameterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnParameterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnParameterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnParameterPropsMixinParameters(props *CfnParameterMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

