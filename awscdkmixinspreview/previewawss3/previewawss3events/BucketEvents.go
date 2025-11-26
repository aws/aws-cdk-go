package previewawss3events

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawss3"
)

// EventBridge event patterns for Bucket.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var bucket Bucket
//
//   bucketEvents := awscdkmixinspreview.BucketEvents_FromBucket(bucket)
//
//   pattern := bucketEvents.ObjectCreatedPattern(&ObjectCreatedProps{
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("us-east-1"),
//   			jsii.String("us-west-2"),
//   		},
//   		Version: []*string{
//   			jsii.String("0"),
//   		},
//   	},
//   })
//
// Experimental.
type BucketEvents interface {
	// EventBridge event pattern for Bucket AWS API Call via CloudTrail.
	// Experimental.
	AwsAPICallViaCloudTrailPattern(options *BucketEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Access Tier Changed.
	// Experimental.
	ObjectAccessTierChangedPattern(options *BucketEvents_ObjectAccessTierChanged_ObjectAccessTierChangedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object ACL Updated.
	// Experimental.
	ObjectACLUpdatedPattern(options *BucketEvents_ObjectACLUpdated_ObjectACLUpdatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Created.
	// Experimental.
	ObjectCreatedPattern(options *BucketEvents_ObjectCreated_ObjectCreatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Deleted.
	// Experimental.
	ObjectDeletedPattern(options *BucketEvents_ObjectDeleted_ObjectDeletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Restore Completed.
	// Experimental.
	ObjectRestoreCompletedPattern(options *BucketEvents_ObjectRestoreCompleted_ObjectRestoreCompletedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Restore Expired.
	// Experimental.
	ObjectRestoreExpiredPattern(options *BucketEvents_ObjectRestoreExpired_ObjectRestoreExpiredProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Restore Initiated.
	// Experimental.
	ObjectRestoreInitiatedPattern(options *BucketEvents_ObjectRestoreInitiated_ObjectRestoreInitiatedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Storage Class Changed.
	// Experimental.
	ObjectStorageClassChangedPattern(options *BucketEvents_ObjectStorageClassChanged_ObjectStorageClassChangedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Tags Added.
	// Experimental.
	ObjectTagsAddedPattern(options *BucketEvents_ObjectTagsAdded_ObjectTagsAddedProps) *awsevents.EventPattern
	// EventBridge event pattern for Bucket Object Tags Deleted.
	// Experimental.
	ObjectTagsDeletedPattern(options *BucketEvents_ObjectTagsDeleted_ObjectTagsDeletedProps) *awsevents.EventPattern
}

// The jsii proxy struct for BucketEvents
type jsiiProxy_BucketEvents struct {
	_ byte // padding
}

// Create BucketEvents from a Bucket reference.
// Experimental.
func BucketEvents_FromBucket(bucketRef interfacesawss3.IBucketRef) BucketEvents {
	_init_.Initialize()

	if err := validateBucketEvents_FromBucketParameters(bucketRef); err != nil {
		panic(err)
	}
	var returns BucketEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents",
		"fromBucket",
		[]interface{}{bucketRef},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) AwsAPICallViaCloudTrailPattern(options *BucketEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps) *awsevents.EventPattern {
	if err := b.validateAwsAPICallViaCloudTrailPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"awsAPICallViaCloudTrailPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectAccessTierChangedPattern(options *BucketEvents_ObjectAccessTierChanged_ObjectAccessTierChangedProps) *awsevents.EventPattern {
	if err := b.validateObjectAccessTierChangedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectAccessTierChangedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectACLUpdatedPattern(options *BucketEvents_ObjectACLUpdated_ObjectACLUpdatedProps) *awsevents.EventPattern {
	if err := b.validateObjectACLUpdatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectACLUpdatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectCreatedPattern(options *BucketEvents_ObjectCreated_ObjectCreatedProps) *awsevents.EventPattern {
	if err := b.validateObjectCreatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectCreatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectDeletedPattern(options *BucketEvents_ObjectDeleted_ObjectDeletedProps) *awsevents.EventPattern {
	if err := b.validateObjectDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectRestoreCompletedPattern(options *BucketEvents_ObjectRestoreCompleted_ObjectRestoreCompletedProps) *awsevents.EventPattern {
	if err := b.validateObjectRestoreCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectRestoreCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectRestoreExpiredPattern(options *BucketEvents_ObjectRestoreExpired_ObjectRestoreExpiredProps) *awsevents.EventPattern {
	if err := b.validateObjectRestoreExpiredPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectRestoreExpiredPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectRestoreInitiatedPattern(options *BucketEvents_ObjectRestoreInitiated_ObjectRestoreInitiatedProps) *awsevents.EventPattern {
	if err := b.validateObjectRestoreInitiatedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectRestoreInitiatedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectStorageClassChangedPattern(options *BucketEvents_ObjectStorageClassChanged_ObjectStorageClassChangedProps) *awsevents.EventPattern {
	if err := b.validateObjectStorageClassChangedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectStorageClassChangedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectTagsAddedPattern(options *BucketEvents_ObjectTagsAdded_ObjectTagsAddedProps) *awsevents.EventPattern {
	if err := b.validateObjectTagsAddedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectTagsAddedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BucketEvents) ObjectTagsDeletedPattern(options *BucketEvents_ObjectTagsDeleted_ObjectTagsDeletedProps) *awsevents.EventPattern {
	if err := b.validateObjectTagsDeletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		b,
		"objectTagsDeletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}

