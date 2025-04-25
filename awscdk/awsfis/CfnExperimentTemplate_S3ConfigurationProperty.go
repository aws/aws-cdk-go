package awsfis


// Specifies the configuration for experiment logging to Amazon S3 .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-s3configuration.html
//
type CfnExperimentTemplate_S3ConfigurationProperty struct {
	// The name of the destination bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-s3configuration.html#cfn-fis-experimenttemplate-s3configuration-bucketname
	//
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fis-experimenttemplate-s3configuration.html#cfn-fis-experimenttemplate-s3configuration-prefix
	//
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

