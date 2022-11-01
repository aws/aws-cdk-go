//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::GameLift
package awscdkgameliftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Script) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Script) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Script) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateScript_FromAssetParameters(scope constructs.Construct, id *string, path *string, options *awss3assets.AssetOptions) error {
	return nil
}

func validateScript_FromBucketParameters(scope constructs.Construct, id *string, bucket awss3.IBucket, key *string) error {
	return nil
}

func validateScript_FromScriptArnParameters(scope constructs.Construct, id *string, scriptArn *string) error {
	return nil
}

func validateScript_FromScriptAttributesParameters(scope constructs.Construct, id *string, attrs *ScriptAttributes) error {
	return nil
}

func validateScript_IsConstructParameters(x interface{}) error {
	return nil
}

func validateScript_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateScript_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewScriptParameters(scope constructs.Construct, id *string, props *ScriptProps) error {
	return nil
}

