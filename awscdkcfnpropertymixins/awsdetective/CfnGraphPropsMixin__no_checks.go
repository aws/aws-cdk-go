//go:build no_runtime_type_checking

package awsdetective

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGraphPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGraphPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGraphPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGraphPropsMixinParameters(props *CfnGraphMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

