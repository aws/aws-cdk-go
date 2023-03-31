package awsstepfunctions


// The `TagsEntry` property specifies *tags* to identify an activity.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsEntryProperty := &tagsEntryProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnActivity_TagsEntryProperty struct {
	// The `key` for a key-value pair in a tag entry.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The `value` for a key-value pair in a tag entry.
	Value *string `field:"required" json:"value" yaml:"value"`
}

