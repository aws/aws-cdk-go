package previewawselasticbeanstalkmixins


// Use the `SourceBundle` property type to specify the Amazon S3 location of the source bundle for an AWS Elastic Beanstalk application version when defining an AWS::ElasticBeanstalk::ApplicationVersion resource in an AWS CloudFormation template.
//
// The `SourceBundle` property is an embedded property of the [AWS::ElasticBeanstalk::ApplicationVersion](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html) resource. It specifies the Amazon S3 location of the source bundle for an AWS Elastic Beanstalk application version.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sourceBundleProperty := &SourceBundleProperty{
//   	S3Bucket: jsii.String("s3Bucket"),
//   	S3Key: jsii.String("s3Key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-sourcebundle.html
//
type CfnApplicationVersionPropsMixin_SourceBundleProperty struct {
	// The Amazon S3 bucket where the data is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-sourcebundle.html#cfn-elasticbeanstalk-applicationversion-sourcebundle-s3bucket
	//
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key where the data is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-applicationversion-sourcebundle.html#cfn-elasticbeanstalk-applicationversion-sourcebundle-s3key
	//
	S3Key *string `field:"optional" json:"s3Key" yaml:"s3Key"`
}

