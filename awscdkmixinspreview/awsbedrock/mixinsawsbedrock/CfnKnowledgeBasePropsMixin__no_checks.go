//go:build no_runtime_type_checking

package mixinsawsbedrock

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnKnowledgeBasePropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnKnowledgeBasePropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnKnowledgeBasePropsMixinParameters(props *CfnKnowledgeBaseMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

