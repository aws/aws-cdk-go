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
//   tagEntryProperty := &TagEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-tagentry.html
//
type CfnEventBus_TagEntryProperty struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-tagentry.html#cfn-events-eventbus-tagentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the specified tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-tagentry.html#cfn-events-eventbus-tagentry-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}

