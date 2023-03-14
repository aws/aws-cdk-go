package awsevents


// A key-value pair associated with an AWS resource.
//
// In EventBridge, rules and event buses support tagging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagEntryProperty := &tagEntryProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnEventBus_TagEntryProperty struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the specified tag key.
	Value *string `field:"required" json:"value" yaml:"value"`
}

