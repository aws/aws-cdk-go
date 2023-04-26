//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BucketBase) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateArnForObjectsParameters(keyPattern *string) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantPutParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	return nil
}

func (b *jsiiProxy_BucketBase) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateBucketBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBucketBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucketBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_BucketBase) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func validateNewBucketBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

