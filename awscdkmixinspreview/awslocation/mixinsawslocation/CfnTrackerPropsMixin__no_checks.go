//go:build no_runtime_type_checking

package mixinsawslocation

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTrackerPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTrackerPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTrackerPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTrackerPropsMixinParameters(props *CfnTrackerMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

