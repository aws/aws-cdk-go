//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnRestApiPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnRestApiPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnRestApiPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnRestApiPropsMixinParameters(props *CfnRestApiMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

