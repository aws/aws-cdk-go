//go:build no_runtime_type_checking

package awsecr

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryBase) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateGrantPullParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateGrantPullPushParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateOnCloudTrailImagePushedParameters(id *string, options *OnCloudTrailImagePushedOptions) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_RepositoryBase) validateOnImageScanCompletedParameters(id *string, options *OnImageScanCompletedOptions) error {
	return nil
}

func validateRepositoryBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRepositoryBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRepositoryBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRepositoryBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

