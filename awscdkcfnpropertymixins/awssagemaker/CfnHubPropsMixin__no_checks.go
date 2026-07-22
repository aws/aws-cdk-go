//go:build no_runtime_type_checking

package awssagemaker

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHubPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHubPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHubPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHubPropsMixinParameters(props *CfnHubMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

