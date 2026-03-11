//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnVPCBlockPublicAccessExclusionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnVPCBlockPublicAccessExclusionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnVPCBlockPublicAccessExclusionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnVPCBlockPublicAccessExclusionPropsMixinParameters(props *CfnVPCBlockPublicAccessExclusionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) error {
	return nil
}

