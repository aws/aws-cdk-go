package mixinsawss3


// Specifies the lifecycle configuration for objects in an Amazon S3 bucket.
//
// For more information, see [Object Lifecycle Management](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon S3 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   				DaysAfterInitiation: jsii.Number(123),
//   			},
//   			ExpirationDate: NewDate(),
//   			ExpirationInDays: jsii.Number(123),
//   			ExpiredObjectDeleteMarker: jsii.Boolean(false),
//   			Id: jsii.String("id"),
//   			NoncurrentVersionExpiration: &NoncurrentVersionExpirationProperty{
//   				NewerNoncurrentVersions: jsii.Number(123),
//   				NoncurrentDays: jsii.Number(123),
//   			},
//   			NoncurrentVersionExpirationInDays: jsii.Number(123),
//   			NoncurrentVersionTransition: &NoncurrentVersionTransitionProperty{
//   				NewerNoncurrentVersions: jsii.Number(123),
//   				StorageClass: jsii.String("storageClass"),
//   				TransitionInDays: jsii.Number(123),
//   			},
//   			NoncurrentVersionTransitions: []interface{}{
//   				&NoncurrentVersionTransitionProperty{
//   					NewerNoncurrentVersions: jsii.Number(123),
//   					StorageClass: jsii.String("storageClass"),
//   					TransitionInDays: jsii.Number(123),
//   				},
//   			},
//   			ObjectSizeGreaterThan: jsii.Number(123),
//   			ObjectSizeLessThan: jsii.Number(123),
//   			Prefix: jsii.String("prefix"),
//   			Status: jsii.String("status"),
//   			TagFilters: []interface{}{
//   				&TagFilterProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			Transition: &TransitionProperty{
//   				StorageClass: jsii.String("storageClass"),
//   				TransitionDate: NewDate(),
//   				TransitionInDays: jsii.Number(123),
//   			},
//   			Transitions: []interface{}{
//   				&TransitionProperty{
//   					StorageClass: jsii.String("storageClass"),
//   					TransitionDate: NewDate(),
//   					TransitionInDays: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	TransitionDefaultMinimumObjectSize: jsii.String("transitionDefaultMinimumObjectSize"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfiguration.html
//
type CfnBucketPropsMixin_LifecycleConfigurationProperty struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfiguration.html#cfn-s3-bucket-lifecycleconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// Indicates which default minimum object size behavior is applied to the lifecycle configuration.
	//
	// > This parameter applies to general purpose buckets only. It isn't supported for directory bucket lifecycle configurations.
	//
	// - `all_storage_classes_128K` - Objects smaller than 128 KB will not transition to any storage class by default.
	// - `varies_by_storage_class` - Objects smaller than 128 KB will transition to Glacier Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes will prevent transitions smaller than 128 KB.
	//
	// To customize the minimum object size for any transition you can add a filter that specifies a custom `ObjectSizeGreaterThan` or `ObjectSizeLessThan` in the body of your transition rule. Custom filters always take precedence over the default transition behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfiguration.html#cfn-s3-bucket-lifecycleconfiguration-transitiondefaultminimumobjectsize
	//
	TransitionDefaultMinimumObjectSize *string `field:"optional" json:"transitionDefaultMinimumObjectSize" yaml:"transitionDefaultMinimumObjectSize"`
}

