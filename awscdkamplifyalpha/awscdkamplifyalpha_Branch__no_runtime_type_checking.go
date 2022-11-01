//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// The CDK Construct Library for AWS::Amplify
package awscdkamplifyalpha

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Branch) validateAddEnvironmentParameters(name *string, value *string) error {
	return nil
}

func (b *jsiiProxy_Branch) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_Branch) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBranch_FromBranchNameParameters(scope constructs.Construct, id *string, branchName *string) error {
	return nil
}

func validateBranch_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBranch_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBranch_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBranchParameters(scope constructs.Construct, id *string, props *BranchProps) error {
	return nil
}

