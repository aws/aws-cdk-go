//go:build no_runtime_type_checking

package awscloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3OriginAccessControl) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_S3OriginAccessControl) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_S3OriginAccessControl) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateS3OriginAccessControl_FromOriginAccessControlIdParameters(scope constructs.Construct, id *string, originAccessControlId *string) error {
	return nil
}

func validateS3OriginAccessControl_IsConstructParameters(x interface{}) error {
	return nil
}

func validateS3OriginAccessControl_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateS3OriginAccessControl_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewS3OriginAccessControlParameters(scope constructs.Construct, id *string, props *S3OriginAccessControlProps) error {
	return nil
}

