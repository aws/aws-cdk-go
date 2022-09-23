package awss3

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Declaration of a Life cycle rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//   var tagFilters interface{}
//
//   lifecycleRule := &lifecycleRule{
//   	abortIncompleteMultipartUploadAfter: cdk.duration.minutes(jsii.Number(30)),
//   	enabled: jsii.Boolean(false),
//   	expiration: cdk.*duration.minutes(jsii.Number(30)),
//   	expirationDate: NewDate(),
//   	expiredObjectDeleteMarker: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	noncurrentVersionExpiration: cdk.*duration.minutes(jsii.Number(30)),
//   	noncurrentVersionsToRetain: jsii.Number(123),
//   	noncurrentVersionTransitions: []noncurrentVersionTransition{
//   		&noncurrentVersionTransition{
//   			storageClass: storageClass,
//   			transitionAfter: cdk.*duration.minutes(jsii.Number(30)),
//
//   			// the properties below are optional
//   			noncurrentVersionsToRetain: jsii.Number(123),
//   		},
//   	},
//   	objectSizeGreaterThan: jsii.Number(123),
//   	objectSizeLessThan: jsii.Number(123),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: map[string]interface{}{
//   		"tagFiltersKey": tagFilters,
//   	},
//   	transitions: []transition{
//   		&transition{
//   			storageClass: storageClass,
//
//   			// the properties below are optional
//   			transitionAfter: cdk.*duration.minutes(jsii.Number(30)),
//   			transitionDate: NewDate(),
//   		},
//   	},
//   }
//
type LifecycleRule struct {
	// Specifies a lifecycle rule that aborts incomplete multipart uploads to an Amazon S3 bucket.
	//
	// The AbortIncompleteMultipartUpload property type creates a lifecycle
	// rule that aborts incomplete multipart uploads to an Amazon S3 bucket.
	// When Amazon S3 aborts a multipart upload, it deletes all parts
	// associated with the multipart upload.
	AbortIncompleteMultipartUploadAfter awscdk.Duration `field:"optional" json:"abortIncompleteMultipartUploadAfter" yaml:"abortIncompleteMultipartUploadAfter"`
	// Whether this rule is enabled.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	Expiration awscdk.Duration `field:"optional" json:"expiration" yaml:"expiration"`
	// Indicates when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	ExpirationDate *time.Time `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions.
	//
	// If set to true, the delete marker will be expired.
	ExpiredObjectDeleteMarker *bool `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
	// A unique identifier for this rule.
	//
	// The value cannot be more than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Time between when a new version of the object is uploaded to the bucket and when old versions of the object expire.
	//
	// For buckets with versioning enabled (or suspended), specifies the time,
	// in days, between when a new version of the object is uploaded to the
	// bucket and when old versions of the object expire. When object versions
	// expire, Amazon S3 permanently deletes them. If you specify a transition
	// and expiration time, the expiration time must be later than the
	// transition time.
	NoncurrentVersionExpiration awscdk.Duration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Indicates a maximum number of noncurrent versions to retain.
	//
	// If there are this many more noncurrent versions,
	// Amazon S3 permanently deletes them.
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
	// One or more transition rules that specify when non-current objects transition to a specified storage class.
	//
	// Only for for buckets with versioning enabled (or suspended).
	//
	// If you specify a transition and expiration time, the expiration time
	// must be later than the transition time.
	NoncurrentVersionTransitions *[]*NoncurrentVersionTransition `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Specifies the minimum object size in bytes for this rule to apply to.
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Specifies the maximum object size in bytes for this rule to apply to.
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The TagFilter property type specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
	TagFilters *map[string]interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// One or more transition rules that specify when an object transitions to a specified storage class.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	Transitions *[]*Transition `field:"optional" json:"transitions" yaml:"transitions"`
}

