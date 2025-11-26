//go:build no_runtime_type_checking

package previewawsdynamodbmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnGlobalTablePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnGlobalTablePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnGlobalTablePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnGlobalTablePropsMixinParameters(props *CfnGlobalTableMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

