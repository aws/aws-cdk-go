//go:build no_runtime_type_checking

package awsecr

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Repository) validateAddLifecycleRuleParameters(rule *LifecycleRule) error {
	return nil
}

func (r *jsiiProxy_Repository) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (r *jsiiProxy_Repository) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGrantPullParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGrantPullPushParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (r *jsiiProxy_Repository) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_Repository) validateOnCloudTrailImagePushedParameters(id *string, options *OnCloudTrailImagePushedOptions) error {
	return nil
}

func (r *jsiiProxy_Repository) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (r *jsiiProxy_Repository) validateOnImageScanCompletedParameters(id *string, options *OnImageScanCompletedOptions) error {
	return nil
}

func validateRepository_ArnForLocalRepositoryParameters(repositoryName *string, scope constructs.IConstruct) error {
	return nil
}

func validateRepository_FromRepositoryArnParameters(scope constructs.Construct, id *string, repositoryArn *string) error {
	return nil
}

func validateRepository_FromRepositoryAttributesParameters(scope constructs.Construct, id *string, attrs *RepositoryAttributes) error {
	return nil
}

func validateRepository_FromRepositoryNameParameters(scope constructs.Construct, id *string, repositoryName *string) error {
	return nil
}

func validateRepository_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRepository_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateRepository_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewRepositoryParameters(scope constructs.Construct, id *string, props *RepositoryProps) error {
	return nil
}

