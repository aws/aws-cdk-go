//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_Bucket) validateAddCorsRuleParameters(rule *CorsRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	if event == "" {
		return fmt.Errorf("parameter event is required, but nil was provided")
	}

	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddInventoryParameters(inventory *Inventory) error {
	if inventory == nil {
		return fmt.Errorf("parameter inventory is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(inventory, func() string { return "parameter inventory" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddLifecycleRuleParameters(rule *LifecycleRule) error {
	if rule == nil {
		return fmt.Errorf("parameter rule is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(rule, func() string { return "parameter rule" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddMetricParameters(metric *BucketMetrics) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(metric, func() string { return "parameter metric" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
	if dest == nil {
		return fmt.Errorf("parameter dest is required, but nil was provided")
	}

	for idx_47e227, v := range *filters {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter filters[%#v]", idx_47e227) }); err != nil {
			return err
		}
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateArnForObjectsParameters(keyPattern *string) error {
	if keyPattern == nil {
		return fmt.Errorf("parameter keyPattern is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	if arnAttr == nil {
		return fmt.Errorf("parameter arnAttr is required, but nil was provided")
	}

	if arnComponents == nil {
		return fmt.Errorf("parameter arnComponents is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(arnComponents, func() string { return "parameter arnComponents" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_Bucket) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBucket_FromBucketArnParameters(scope constructs.Construct, id *string, bucketArn *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if bucketArn == nil {
		return fmt.Errorf("parameter bucketArn is required, but nil was provided")
	}

	return nil
}

func validateBucket_FromBucketAttributesParameters(scope constructs.Construct, id *string, attrs *BucketAttributes) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if attrs == nil {
		return fmt.Errorf("parameter attrs is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(attrs, func() string { return "parameter attrs" }); err != nil {
		return err
	}

	return nil
}

func validateBucket_FromBucketNameParameters(scope constructs.Construct, id *string, bucketName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if bucketName == nil {
		return fmt.Errorf("parameter bucketName is required, but nil was provided")
	}

	return nil
}

func validateBucket_FromCfnBucketParameters(cfnBucket CfnBucket) error {
	if cfnBucket == nil {
		return fmt.Errorf("parameter cfnBucket is required, but nil was provided")
	}

	return nil
}

func validateBucket_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateBucket_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBucket_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBucket_ValidateBucketNameParameters(physicalName *string) error {
	if physicalName == nil {
		return fmt.Errorf("parameter physicalName is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Bucket) validateSetAutoCreatePolicyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBucketParameters(scope constructs.Construct, id *string, props *BucketProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

