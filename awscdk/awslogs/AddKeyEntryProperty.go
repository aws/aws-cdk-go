package awslogs


// This object defines one key that will be added with the addKeys processor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addKeyEntryProperty := &AddKeyEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//
//   	// the properties below are optional
//   	OverwriteIfExists: jsii.Boolean(false),
//   }
//
type AddKeyEntryProperty struct {
	// The key of the new entry to be added to the log event.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the new entry to be added to the log event.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Specifies whether to overwrite the value if the key already exists.
	// Default: false.
	//
	OverwriteIfExists *bool `field:"optional" json:"overwriteIfExists" yaml:"overwriteIfExists"`
}

