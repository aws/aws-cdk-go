//go:build no_runtime_type_checking

package awscognito

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTermsPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTermsPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTermsPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTermsPropsMixinParameters(props *CfnTermsMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

