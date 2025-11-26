//go:build no_runtime_type_checking

package previewawsopensearchserverlessmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnCollectionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnCollectionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnCollectionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnCollectionPropsMixinParameters(props *CfnCollectionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

