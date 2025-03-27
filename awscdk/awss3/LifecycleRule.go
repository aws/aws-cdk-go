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
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var storageClass storageClass
//   var tagFilters interface{}
//
//   lifecycleRule := &LifecycleRule{
//   	AbortIncompleteMultipartUploadAfter: cdk.Duration_Minutes(jsii.Number(30)),
//   	Enabled: jsii.Boolean(false),
//   	Expiration: cdk.Duration_*Minutes(jsii.Number(30)),
//   	ExpirationDate: NewDate(),
//   	ExpiredObjectDeleteMarker: jsii.Boolean(false),
//   	Id: jsii.String("id"),
//   	NoncurrentVersionExpiration: cdk.Duration_*Minutes(jsii.Number(30)),
//   	NoncurrentVersionsToRetain: jsii.Number(123),
//   	NoncurrentVersionTransitions: []noncurrentVersionTransition{
//   		&noncurrentVersionTransition{
//   			StorageClass: storageClass,
//   			TransitionAfter: cdk.Duration_*Minutes(jsii.Number(30)),
//
//   			// the properties below are optional
//   			NoncurrentVersionsToRetain: jsii.Number(123),
//   		},
//   	},
//   	ObjectSizeGreaterThan: jsii.Number(123),
//   	ObjectSizeLessThan: jsii.Number(123),
//   	Prefix: jsii.String("prefix"),
//   	TagFilters: map[string]interface{}{
//   		"tagFiltersKey": tagFilters,
//   	},
//   	Transitions: []transition{
//   		&transition{
//   			StorageClass: storageClass,
//
//   			// the properties below are optional
//   			TransitionAfter: cdk.Duration_*Minutes(jsii.Number(30)),
//   			TransitionDate: NewDate(),
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
	//
	// The underlying configuration is expressed in whole numbers of days. Providing a Duration that
	// does not represent a whole number of days will result in a runtime or deployment error.
	// Default: - Incomplete uploads are never aborted.
	//
	AbortIncompleteMultipartUploadAfter awscdk.Duration `field:"optional" json:"abortIncompleteMultipartUploadAfter" yaml:"abortIncompleteMultipartUploadAfter"`
	// Whether this rule is enabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	//
	// The underlying configuration is expressed in whole numbers of days. Providing a Duration that
	// does not represent a whole number of days will result in a runtime or deployment error.
	// Default: - No expiration timeout.
	//
	Expiration awscdk.Duration `field:"optional" json:"expiration" yaml:"expiration"`
	// Indicates when objects are deleted from Amazon S3 and Amazon Glacier.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	// Default: - No expiration date.
	//
	ExpirationDate *time.Time `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions.
	//
	// If set to true, the delete marker will be expired.
	// Default: false.
	//
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
	//
	// The underlying configuration is expressed in whole numbers of days. Providing a Duration that
	// does not represent a whole number of days will result in a runtime or deployment error.
	// Default: - No noncurrent version expiration.
	//
	NoncurrentVersionExpiration awscdk.Duration `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// Indicates a maximum number of noncurrent versions to retain.
	//
	// If there are this many more noncurrent versions,
	// Amazon S3 permanently deletes them.
	// Default: - No noncurrent versions to retain.
	//
	NoncurrentVersionsToRetain *float64 `field:"optional" json:"noncurrentVersionsToRetain" yaml:"noncurrentVersionsToRetain"`
	// One or more transition rules that specify when non-current objects transition to a specified storage class.
	//
	// Only for for buckets with versioning enabled (or suspended).
	//
	// If you specify a transition and expiration time, the expiration time
	// must be later than the transition time.
	NoncurrentVersionTransitions *[]*NoncurrentVersionTransition `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Specifies the minimum object size in bytes for this rule to apply to.
	//
	// Objects must be larger than this value in bytes.
	// Default: - No rule.
	//
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Specifies the maximum object size in bytes for this rule to apply to.
	//
	// Objects must be smaller than this value in bytes.
	// Default: - No rule.
	//
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	// Default: - Rule applies to all objects.
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The TagFilter property type specifies tags to use to identify a subset of objects for an Amazon S3 bucket.
	// Default: - Rule applies to all objects.
	//
	TagFilters *map[string]interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// One or more transition rules that specify when an object transitions to a specified storage class.
	//
	// If you specify an expiration and transition time, you must use the same
	// time unit for both properties (either in days or by date). The
	// expiration time must also be later than the transition time.
	// Default: - No transition rules.
	//
	Transitions *[]*Transition `field:"optional" json:"transitions" yaml:"transitions"`
}

