//go:build no_runtime_type_checking

package previewawsemrmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWALWorkspacePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWALWorkspacePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWALWorkspacePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWALWorkspacePropsMixinParameters(props *CfnWALWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

