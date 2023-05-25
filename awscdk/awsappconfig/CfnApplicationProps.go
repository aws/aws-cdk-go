package awsappconfig


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &CfnApplicationProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []tagsProperty{
//   		&tagsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnApplicationProps struct {
	// A name for the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the application.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata to assign to the application.
	//
	// Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*CfnApplication_TagsProperty `field:"optional" json:"tags" yaml:"tags"`
}

