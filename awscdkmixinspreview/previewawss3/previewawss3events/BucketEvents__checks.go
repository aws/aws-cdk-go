//go:build !no_runtime_type_checking

package previewawss3events

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

func (b *jsiiProxy_BucketEvents) validateAwsAPICallViaCloudTrailPatternParameters(options *AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectAccessTierChangedPatternParameters(options *ObjectAccessTierChanged_ObjectAccessTierChangedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectACLUpdatedPatternParameters(options *ObjectACLUpdated_ObjectACLUpdatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectCreatedPatternParameters(options *ObjectCreated_ObjectCreatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectDeletedPatternParameters(options *ObjectDeleted_ObjectDeletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreCompletedPatternParameters(options *ObjectRestoreCompleted_ObjectRestoreCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreExpiredPatternParameters(options *ObjectRestoreExpired_ObjectRestoreExpiredProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectRestoreInitiatedPatternParameters(options *ObjectRestoreInitiated_ObjectRestoreInitiatedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectStorageClassChangedPatternParameters(options *ObjectStorageClassChanged_ObjectStorageClassChangedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectTagsAddedPatternParameters(options *ObjectTagsAdded_ObjectTagsAddedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BucketEvents) validateObjectTagsDeletedPatternParameters(options *ObjectTagsDeleted_ObjectTagsDeletedProps) error {
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

