//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnDaemonTaskDefinitionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnDaemonTaskDefinitionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnDaemonTaskDefinitionPropsMixinParameters(props *CfnDaemonTaskDefinitionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

