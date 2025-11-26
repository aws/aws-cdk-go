//go:build !no_runtime_type_checking

package previewawss3events

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (b *jsiiProxy_BucketEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *BucketEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectAccessTierChangedPatternParameters(options *BucketEvents_ObjectAccessTierChanged_ObjectAccessTierChangedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectACLUpdatedPatternParameters(options *BucketEvents_ObjectACLUpdated_ObjectACLUpdatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectCreatedPatternParameters(options *BucketEvents_ObjectCreated_ObjectCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectDeletedPatternParameters(options *BucketEvents_ObjectDeleted_ObjectDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreCompletedPatternParameters(options *BucketEvents_ObjectRestoreCompleted_ObjectRestoreCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreExpiredPatternParameters(options *BucketEvents_ObjectRestoreExpired_ObjectRestoreExpiredProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreInitiatedPatternParameters(options *BucketEvents_ObjectRestoreInitiated_ObjectRestoreInitiatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectStorageClassChangedPatternParameters(options *BucketEvents_ObjectStorageClassChanged_ObjectStorageClassChangedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectTagsAddedPatternParameters(options *BucketEvents_ObjectTagsAdded_ObjectTagsAddedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectTagsDeletedPatternParameters(options *BucketEvents_ObjectTagsDeleted_ObjectTagsDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateBucketEvents_FromBucketParameters(bucketRef interfacesawss3.IBucketRef) error {
	if bucketRef == nil {
		return fmt.Errorf("parameter bucketRef is required, but nil was provided")
	}

	return nil
}

