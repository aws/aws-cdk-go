//go:build no_runtime_type_checking

package awsecr

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRepository) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantPullParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateGrantPullPushParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnCloudTrailEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnCloudTrailImagePushedParameters(id *string, options *OnCloudTrailImagePushedOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnEventParameters(id *string, options *awsevents.OnEventOptions) error {
	return nil
}

func (i *jsiiProxy_IRepository) validateOnImageScanCompletedParameters(id *string, options *OnImageScanCompletedOptions) error {
	return nil
}

