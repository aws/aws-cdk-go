//go:build no_runtime_type_checking

package previewawstimestreammixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnInfluxDBClusterPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnInfluxDBClusterPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnInfluxDBClusterPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnInfluxDBClusterPropsMixinParameters(props *CfnInfluxDBClusterMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

