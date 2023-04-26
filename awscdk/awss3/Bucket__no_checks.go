//go:build no_runtime_type_checking

package awss3

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Bucket) validateAddCorsRuleParameters(rule *CorsRule) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddInventoryParameters(inventory *Inventory) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddLifecycleRuleParameters(rule *LifecycleRule) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddMetricParameters(metric *BucketMetrics) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateArnForObjectsParameters(keyPattern *string) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	return nil
}

func (b *jsiiProxy_Bucket) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	return nil
}

func validateBucket_FromBucketArnParameters(scope constructs.Construct, id *string, bucketArn *string) error {
	return nil
}

func validateBucket_FromBucketAttributesParameters(scope constructs.Construct, id *string, attrs *BucketAttributes) error {
	return nil
}

func validateBucket_FromBucketNameParameters(scope constructs.Construct, id *string, bucketName *string) error {
	return nil
}

func validateBucket_FromCfnBucketParameters(cfnBucket CfnBucket) error {
	return nil
}

func validateBucket_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucket_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBucket_ValidateBucketNameParameters(physicalName *string) error {
	return nil
}

func (j *jsiiProxy_Bucket) validateSetAutoCreatePolicyParameters(val *bool) error {
	return nil
}

func validateNewBucketParameters(scope constructs.Construct, id *string, props *BucketProps) error {
	return nil
}

