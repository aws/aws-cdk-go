//go:build no_runtime_type_checking

package previewawss3expressmixins

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) validateApplyToParameters(construct constructs.IConstruct) error {
	return nil
}

func (c *jsiiProxy_CfnBucketPolicyPropsMixin) validateSupportsParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCfnBucketPolicyPropsMixin_IsMixinParameters(x interface{}) error {
	return nil
}

func validateNewCfnBucketPolicyPropsMixinParameters(props *CfnBucketPolicyMixinProps, options *mixins.CfnPropertyMixinOptions) error {
	return nil
}

