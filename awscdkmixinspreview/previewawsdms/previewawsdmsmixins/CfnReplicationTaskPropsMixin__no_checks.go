//go:build no_runtime_type_checking

package previewawsdmsmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnReplicationTaskPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnReplicationTaskPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnReplicationTaskPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnReplicationTaskPropsMixinParameters(props *CfnReplicationTaskMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

