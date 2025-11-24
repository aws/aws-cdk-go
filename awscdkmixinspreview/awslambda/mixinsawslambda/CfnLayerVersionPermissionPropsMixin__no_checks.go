//go:build no_runtime_type_checking

package mixinsawslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnLayerVersionPermissionPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnLayerVersionPermissionPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnLayerVersionPermissionPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnLayerVersionPermissionPropsMixinParameters(props *CfnLayerVersionPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

