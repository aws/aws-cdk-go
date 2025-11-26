package previewawss3outpostsmixins


// The container for the lifecycle configuration for the objects stored in an S3 on Outposts bucket.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var filter interface{}
//
//   lifecycleConfigurationProperty := &LifecycleConfigurationProperty{
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   				DaysAfterInitiation: jsii.Number(123),
//   			},
//   			ExpirationDate: jsii.String("expirationDate"),
//   			ExpirationInDays: jsii.Number(123),
//   			Filter: filter,
//   			Id: jsii.String("id"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
//
type CfnBucketPropsMixin_LifecycleConfigurationProperty struct {
	// The container for the lifecycle configuration rules for the objects stored in the S3 on Outposts bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

