package awsservicecatalog


// Properties for defining a `CfnTagOption`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagOptionProps := &cfnTagOptionProps{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//
//   	// the properties below are optional
//   	active: jsii.Boolean(false),
//   }
//
type CfnTagOptionProps struct {
	// The TagOption key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The TagOption value.
	Value *string `field:"required" json:"value" yaml:"value"`
	// The TagOption active state.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
}

