//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BucketBase) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (b *jsiiProxy_BucketBase) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (b *jsiiProxy_BucketBase) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (b *jsiiProxy_BucketBase) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	if policy == "" {
		return fmt.Errorf("parameter policy is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateArnForObjectsParameters(keyPattern *string) error {
	if keyPattern == nil {
		return fmt.Errorf("parameter keyPattern is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
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

func (b *jsiiProxy_BucketBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	if nameAttr == nil {
		return fmt.Errorf("parameter nameAttr is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantPutParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantReadParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketBase) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBucketBase_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateBucketBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateBucketBase_IsResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BucketBase) validateSetAutoCreatePolicyParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBucketBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
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

