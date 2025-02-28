package awss3express


// Container for lifecycle rules. You can add as many as 1000 rules.
//
// For more information see, [Creating and managing a lifecycle configuration for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-objects-lifecycle.html          ) in the *Amazon S3 User Guide* .
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
//   			ExpirationInDays: jsii.Number(123),
//   			Id: jsii.String("id"),
//   			ObjectSizeGreaterThan: jsii.String("objectSizeGreaterThan"),
//   			ObjectSizeLessThan: jsii.String("objectSizeLessThan"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-lifecycleconfiguration.html
//
type CfnDirectoryBucket_LifecycleConfigurationProperty struct {
	// A lifecycle rule for individual objects in an Amazon S3 Express bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3express-directorybucket-lifecycleconfiguration.html#cfn-s3express-directorybucket-lifecycleconfiguration-rules
	//
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

