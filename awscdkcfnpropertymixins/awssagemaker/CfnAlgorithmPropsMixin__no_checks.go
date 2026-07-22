//go:build no_runtime_type_checking

package awssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnAlgorithmPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnAlgorithmPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnAlgorithmPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnAlgorithmPropsMixinParameters(props *CfnAlgorithmMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

