package awss3


// Specifies lifecycle rules for an Amazon S3 bucket.
//
// For more information, see [Put Bucket Lifecycle Configuration](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlifecycle.html) in the *Amazon S3 API Reference* .
//
// You must specify at least one of the following properties: `AbortIncompleteMultipartUpload` , `ExpirationDate` , `ExpirationInDays` , `NoncurrentVersionExpirationInDays` , `NoncurrentVersionTransition` , `NoncurrentVersionTransitions` , `Transition` , or `Transitions` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleProperty := &ruleProperty{
//   	status: jsii.String("status"),
//
//   	// the properties below are optional
//   	abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   		daysAfterInitiation: jsii.Number(123),
//   	},
//   	expirationDate: NewDate(),
//   	expirationInDays: jsii.Number(123),
//   	expiredObjectDeleteMarker: jsii.Boolean(false),
//   	id: jsii.String("id"),
//   	noncurrentVersionExpiration: &noncurrentVersionExpirationProperty{
//   		noncurrentDays: jsii.Number(123),
//
//   		// the properties below are optional
//   		newerNoncurrentVersions: jsii.Number(123),
//   	},
//   	noncurrentVersionExpirationInDays: jsii.Number(123),
//   	noncurrentVersionTransition: &noncurrentVersionTransitionProperty{
//   		storageClass: jsii.String("storageClass"),
//   		transitionInDays: jsii.Number(123),
//
//   		// the properties below are optional
//   		newerNoncurrentVersions: jsii.Number(123),
//   	},
//   	noncurrentVersionTransitions: []interface{}{
//   		&noncurrentVersionTransitionProperty{
//   			storageClass: jsii.String("storageClass"),
//   			transitionInDays: jsii.Number(123),
//
//   			// the properties below are optional
//   			newerNoncurrentVersions: jsii.Number(123),
//   		},
//   	},
//   	objectSizeGreaterThan: jsii.Number(123),
//   	objectSizeLessThan: jsii.Number(123),
//   	prefix: jsii.String("prefix"),
//   	tagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	transition: &transitionProperty{
//   		storageClass: jsii.String("storageClass"),
//
//   		// the properties below are optional
//   		transitionDate: NewDate(),
//   		transitionInDays: jsii.Number(123),
//   	},
//   	transitions: []interface{}{
//   		&transitionProperty{
//   			storageClass: jsii.String("storageClass"),
//
//   			// the properties below are optional
//   			transitionDate: NewDate(),
//   			transitionInDays: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnBucket_RuleProperty struct {
	// If `Enabled` , the rule is currently being applied.
	//
	// If `Disabled` , the rule is not currently being applied.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Specifies a lifecycle rule that stops incomplete multipart uploads to an Amazon S3 bucket.
	AbortIncompleteMultipartUpload interface{} `field:"optional" json:"abortIncompleteMultipartUpload" yaml:"abortIncompleteMultipartUpload"`
	// Indicates when objects are deleted from Amazon S3 and Amazon S3 Glacier.
	//
	// The date value must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
	ExpirationDate interface{} `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon S3 Glacier.
	//
	// If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.
	ExpirationInDays *float64 `field:"optional" json:"expirationInDays" yaml:"expirationInDays"`
	// Indicates whether Amazon S3 will remove a delete marker without any noncurrent versions.
	//
	// If set to true, the delete marker will be removed if there are no noncurrent versions. This cannot be specified with `ExpirationInDays` , `ExpirationDate` , or `TagFilters` .
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
	// Unique identifier for the rule.
	//
	// The value can't be longer than 255 characters.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Specifies when noncurrent object versions expire.
	//
	// Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.
	NoncurrentVersionExpiration interface{} `field:"optional" json:"noncurrentVersionExpiration" yaml:"noncurrentVersionExpiration"`
	// (Deprecated.) For buckets with versioning enabled (or suspended), specifies the time, in days, between when a new version of the object is uploaded to the bucket and when old versions of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you specify a transition and expiration time, the expiration time must be later than the transition time.
	NoncurrentVersionExpirationInDays *float64 `field:"optional" json:"noncurrentVersionExpirationInDays" yaml:"noncurrentVersionExpirationInDays"`
	// (Deprecated.) For buckets with versioning enabled (or suspended), specifies when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransitions` property.
	NoncurrentVersionTransition interface{} `field:"optional" json:"noncurrentVersionTransition" yaml:"noncurrentVersionTransition"`
	// For buckets with versioning enabled (or suspended), one or more transition rules that specify when non-current objects transition to a specified storage class.
	//
	// If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransition` property.
	NoncurrentVersionTransitions interface{} `field:"optional" json:"noncurrentVersionTransitions" yaml:"noncurrentVersionTransitions"`
	// Specifies the minimum object size in bytes for this rule to apply to.
	//
	// Objects must be larger than this value in bytes. For more information about size based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide* .
	ObjectSizeGreaterThan *float64 `field:"optional" json:"objectSizeGreaterThan" yaml:"objectSizeGreaterThan"`
	// Specifies the maximum object size in bytes for this rule to apply to.
	//
	// Objects must be smaller than this value in bytes. For more information about sized based rules, see [Lifecycle configuration using size-based rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/lifecycle-configuration-examples.html#lc-size-rules) in the *Amazon S3 User Guide* .
	ObjectSizeLessThan *float64 `field:"optional" json:"objectSizeLessThan" yaml:"objectSizeLessThan"`
	// Object key prefix that identifies one or more objects to which this rule applies.
	//
	// > Replacement must be made for object keys containing special characters (such as carriage returns) when using XML requests. For more information, see [XML related object key constraints](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-keys.html#object-key-xml-related-constraints) .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Tags to use to identify a subset of objects to which the lifecycle rule applies.
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// (Deprecated.) Specifies when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transitions` property.
	Transition interface{} `field:"optional" json:"transition" yaml:"transition"`
	// One or more transition rules that specify when an object transitions to a specified storage class.
	//
	// If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transition` property.
	Transitions interface{} `field:"optional" json:"transitions" yaml:"transitions"`
}

