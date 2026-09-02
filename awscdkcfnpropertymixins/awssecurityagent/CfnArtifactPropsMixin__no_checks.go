//go:build no_runtime_type_checking

package awssecurityagent

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnArtifactPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnArtifactPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnArtifactPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnArtifactPropsMixinParameters(props *CfnArtifactMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

