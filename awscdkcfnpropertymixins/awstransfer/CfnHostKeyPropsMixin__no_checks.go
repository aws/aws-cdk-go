//go:build no_runtime_type_checking

package awstransfer

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnHostKeyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnHostKeyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnHostKeyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnHostKeyPropsMixinParameters(props *CfnHostKeyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

