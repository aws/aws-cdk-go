package awselasticbeanstalk


// Properties for defining a `CfnApplicationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationVersionProps := &CfnApplicationVersionProps{
//   	ApplicationName: jsii.String("applicationName"),
//   	SourceBundle: &SourceBundleProperty{
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html
//
type CfnApplicationVersionProps struct {
	// The name of the Elastic Beanstalk application that is associated with this application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-applicationname
	//
	ApplicationName interface{} `field:"required" json:"applicationName" yaml:"applicationName"`
	// The Amazon S3 bucket and key that identify the location of the source bundle for this version.
	//
	// > The Amazon S3 bucket must be in the same region as the environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	//
	SourceBundle interface{} `field:"required" json:"sourceBundle" yaml:"sourceBundle"`
	// A description of this application version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-applicationversion.html#cfn-elasticbeanstalk-applicationversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

