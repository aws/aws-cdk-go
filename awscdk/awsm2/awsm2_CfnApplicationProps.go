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
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnApplicationProps struct {
	// `AWS::M2::Application.Definition`.
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// `AWS::M2::Application.EngineType`.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// `AWS::M2::Application.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::M2::Application.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::M2::Application.KmsKeyId`.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// `AWS::M2::Application.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

