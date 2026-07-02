//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMicrovmImagePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMicrovmImagePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMicrovmImagePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMicrovmImagePropsMixinParameters(props *CfnMicrovmImageMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

