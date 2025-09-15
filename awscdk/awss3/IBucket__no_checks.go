//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IBucket) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateAddReplicationPolicyParameters(roleArn *string) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateArnForObjectsParameters(keyPattern *string) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantReplicationPermissionParameters(identity awsiam.IGrantable, props *GrantReplicationPermissionProps) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	return nil
}

func (i *jsiiProxy_IBucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

