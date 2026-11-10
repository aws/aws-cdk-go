//go:build no_runtime_type_checking

package awsmgn

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWavePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWavePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWavePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWavePropsMixinParameters(props *CfnWaveMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

