package awsstepfunctionstasks


// The key/value pair for a tag.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customModelTag := &CustomModelTag{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CustomModelTag struct {
	// Key for the tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value for the tag.
	Value *string `field:"required" json:"value" yaml:"value"`
}

