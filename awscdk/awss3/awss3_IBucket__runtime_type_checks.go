//go:build !no_runtime_type_checking

package awss3

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

func (i *jsiiProxy_IBucket) validateAddEventNotificationParameters(event EventType, dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (i *jsiiProxy_IBucket) validateAddObjectCreatedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (i *jsiiProxy_IBucket) validateAddObjectRemovedNotificationParameters(dest IBucketNotificationDestination, filters *[]*NotificationKeyFilter) error {
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

func (i *jsiiProxy_IBucket) validateAddToResourcePolicyParameters(permission awsiam.PolicyStatement) error {
	if permission == nil {
		return fmt.Errorf("parameter permission is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateArnForObjectsParameters(keyPattern *string) error {
	if keyPattern == nil {
		return fmt.Errorf("parameter keyPattern is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantDeleteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantPutParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantPutAclParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantReadParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantReadWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateGrantWriteParameters(identity awsiam.IGrantable) error {
	if identity == nil {
		return fmt.Errorf("parameter identity is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailEventParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailPutObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateOnCloudTrailWriteObjectParameters(id *string, options *OnCloudTrailBucketEventOptions) error {
	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateTransferAccelerationUrlForObjectParameters(options *TransferAccelerationUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (i *jsiiProxy_IBucket) validateVirtualHostedUrlForObjectParameters(options *VirtualHostedStyleUrlOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

