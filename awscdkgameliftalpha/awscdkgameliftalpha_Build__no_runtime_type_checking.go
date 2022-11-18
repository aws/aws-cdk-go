//go:build no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Build) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_Build) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_Build) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBuild_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateBuild_FromBucketParameters(scope constructs.Construct, id *string, bucket awss3.IBucket, key *string) error {
	return nil
}

func validateBuild_FromBuildAttributesParameters(scope constructs.Construct, id *string, attrs *BuildAttributes) error {
	return nil
}

func validateBuild_FromBuildIdParameters(scope constructs.Construct, id *string, buildId *string) error {
	return nil
}

func validateBuild_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBuild_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBuild_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBuildParameters(scope constructs.Construct, id *string, props *BuildProps) error {
	return nil
}

