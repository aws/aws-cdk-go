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
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   				DaysAfterInitiation: jsii.Number(123),
//   			},
//   			ExpirationDate: NewDate(),
//   			ExpirationInDays: jsii.Number(123),
//   			ExpiredObjectDeleteMarker: jsii.Boolean(false),
//   			Id: jsii.String("id"),
//   			NoncurrentVersionExpiration: &NoncurrentVersionExpirationProperty{
//   				NoncurrentDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				NewerNoncurrentVersions: jsii.Number(123),
//   			},
//   			NoncurrentVersionExpirationInDays: jsii.Number(123),
//   			NoncurrentVersionTransition: &NoncurrentVersionTransitionProperty{
//   				StorageClass: jsii.String("storageClass"),
//   				TransitionInDays: jsii.Number(123),
//
//   				// the properties below are optional
//   				NewerNoncurrentVersions: jsii.Number(123),
//   			},
//   			NoncurrentVersionTransitions: []interface{}{
//   				&NoncurrentVersionTransitionProperty{
//   					StorageClass: jsii.String("storageClass"),
//   					TransitionInDays: jsii.Number(123),
//
//   					// the properties below are optional
//   					NewerNoncurrentVersions: jsii.Number(123),
//   				},
//   			},
//   			ObjectSizeGreaterThan: jsii.Number(123),
//   			ObjectSizeLessThan: jsii.Number(123),
//   			Prefix: jsii.String("prefix"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Transition: &TransitionProperty{
//   				StorageClass: jsii.String("storageClass"),
//
//   				// the properties below are optional
//   				TransitionDate: NewDate(),
//   				TransitionInDays: jsii.Number(123),
//   			},
//   			Transitions: []interface{}{
//   				&TransitionProperty{
//   					StorageClass: jsii.String("storageClass"),
//
//   					// the properties below are optional
//   					TransitionDate: NewDate(),
//   					TransitionInDays: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfiguration.html
//
type CfnBucket_LifecycleConfigurationProperty struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfiguration.html#cfn-s3-bucket-lifecycleconfiguration-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

