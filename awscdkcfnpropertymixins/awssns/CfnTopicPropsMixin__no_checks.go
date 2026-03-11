//go:build no_runtime_type_checking

package awssns

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnTopicPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnTopicPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnTopicPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnTopicPropsMixinParameters(props *CfnTopicMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

