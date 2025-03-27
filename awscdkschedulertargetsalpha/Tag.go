package awscdkschedulertargetsalpha


// Metadata that you apply to a resource to help categorize and organize the resource.
//
// Each tag consists of a key and an optional value, both of which you define.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import scheduler_targets_alpha "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha"
//
//   tag := &Tag{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// Deprecated.
type Tag struct {
	// Key is the name of the tag.
	// Deprecated.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Value is the metadata contents of the tag.
	// Deprecated.
	Value *string `field:"required" json:"value" yaml:"value"`
}

