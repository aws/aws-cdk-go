package awseventstargets


// Metadata that you apply to a resource to help categorize and organize the resource.
//
// Each tag consists of a key and an optional value, both of which you define.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tag := &Tag{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type Tag struct {
	// Key is the name of the tag.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value is the metadata contents of the tag.
	Value *string `field:"required" json:"value" yaml:"value"`
}

