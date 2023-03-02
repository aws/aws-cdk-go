package awsservicecatalogappregistry


// Properties for defining a `CfnApplication`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApplicationProps := &cfnApplicationProps{
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
	// The name of the application.
	//
	// The name must be unique in the region in which you are creating the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the application.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Key-value pairs you can use to associate with the application.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

