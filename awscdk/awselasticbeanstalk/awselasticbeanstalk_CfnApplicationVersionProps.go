package awselasticbeanstalk


// Properties for defining a `CfnApplicationVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationVersionProps := &cfnApplicationVersionProps{
//   	applicationName: jsii.String("applicationName"),
//   	sourceBundle: &sourceBundleProperty{
//   		s3Bucket: jsii.String("s3Bucket"),
//   		s3Key: jsii.String("s3Key"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   }
//
type CfnApplicationVersionProps struct {
	// The name of the Elastic Beanstalk application that is associated with this application version.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// The Amazon S3 bucket and key that identify the location of the source bundle for this version.
	//
	// > The Amazon S3 bucket must be in the same region as the environment.
	SourceBundle interface{} `field:"required" json:"sourceBundle" yaml:"sourceBundle"`
	// A description of this application version.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

