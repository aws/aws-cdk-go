package awspinpoint


// Properties for defining a `CfnApp`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnAppProps := &cfnAppProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	tags: tags,
//   }
//
type CfnAppProps struct {
	// The display name of the application.
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

