//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketPolicy) validateApplyRemovalPolicyParameters(removalPolicy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BucketPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BucketPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBucketPolicy_FromCfnBucketPolicyParameters(cfnBucketPolicy CfnBucketPolicy) error {
	return nil
}

func validateBucketPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBucketPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucketPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBucketPolicyParameters(scope constructs.Construct, id *string, props *BucketPolicyProps) error {
	return nil
}

