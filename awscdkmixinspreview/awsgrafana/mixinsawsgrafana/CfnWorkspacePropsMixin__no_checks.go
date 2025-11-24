//go:build no_runtime_type_checking

package mixinsawsgrafana

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnWorkspacePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnWorkspacePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnWorkspacePropsMixinParameters(props *CfnWorkspaceMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

