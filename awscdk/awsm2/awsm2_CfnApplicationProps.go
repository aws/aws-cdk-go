package awsm2


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
//   	definition: &definitionProperty{
//   		content: jsii.String("content"),
//   		s3Location: jsii.String("s3Location"),
//   	},
//   	engineType: jsii.String("engineType"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnApplicationProps struct {
	// The application definition for a particular application. You can specify either inline JSON or an Amazon S3 bucket location.
	//
	// For information about application definitions, see the [AWS Mainframe Modernization User Guide](https://docs.aws.amazon.com/m2/latest/userguide/applications-m2-definition.html) .
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// The type of the target platform for this application.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The name of the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the application.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

