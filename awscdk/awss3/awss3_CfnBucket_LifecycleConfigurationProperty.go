package awss3


// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecycleConfigurationProperty := &lifecycleConfigurationProperty{
//   	rules: []interface{}{
//   		&ruleProperty{
//   			status: jsii.String("status"),
//
//   			// the properties below are optional
//   			abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   				daysAfterInitiation: jsii.Number(123),
//   			},
//   			expirationDate: NewDate(),
//   			expirationInDays: jsii.Number(123),
//   			expiredObjectDeleteMarker: jsii.Boolean(false),
//   			id: jsii.String("id"),
//   			noncurrentVersionExpiration: &noncurrentVersionExpirationProperty{
//   				noncurrentDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				newerNoncurrentVersions: jsii.Number(123),
//   			},
//   			noncurrentVersionExpirationInDays: jsii.Number(123),
//   			noncurrentVersionTransition: &noncurrentVersionTransitionProperty{
//   				storageClass: jsii.String("storageClass"),
//   				transitionInDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				newerNoncurrentVersions: jsii.Number(123),
//   			},
//   			noncurrentVersionTransitions: []interface{}{
//   				&noncurrentVersionTransitionProperty{
//   					storageClass: jsii.String("storageClass"),
//   					transitionInDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					newerNoncurrentVersions: jsii.Number(123),
//   				},
//   			},
//   			objectSizeGreaterThan: jsii.Number(123),
//   			objectSizeLessThan: jsii.Number(123),
//   			prefix: jsii.String("prefix"),
//   			tagFilters: []interface{}{
//   				&tagFilterProperty{
//   					key: jsii.String("key"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			transition: &transitionProperty{
//   				storageClass: jsii.String("storageClass"),
//
//   				// the properties below are optional
//   				transitionDate: NewDate(),
//   				transitionInDays: jsii.Number(123),
//   			},
//   			transitions: []interface{}{
//   				&transitionProperty{
//   					storageClass: jsii.String("storageClass"),
//
//   					// the properties below are optional
//   					transitionDate: NewDate(),
//   					transitionInDays: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnBucket_LifecycleConfigurationProperty struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

