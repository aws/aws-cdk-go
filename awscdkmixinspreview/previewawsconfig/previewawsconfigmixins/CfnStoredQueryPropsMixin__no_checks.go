//go:build no_runtime_type_checking

package previewawsconfigmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnStoredQueryPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnStoredQueryPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnStoredQueryPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnStoredQueryPropsMixinParameters(props *CfnStoredQueryMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

