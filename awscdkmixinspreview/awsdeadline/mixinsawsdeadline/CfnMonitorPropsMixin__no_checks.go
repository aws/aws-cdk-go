//go:build no_runtime_type_checking

package mixinsawsdeadline

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnMonitorPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnMonitorPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnMonitorPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnMonitorPropsMixinParameters(props *CfnMonitorMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

