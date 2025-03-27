package awsevents


// A key-value pair associated with an ECS Target of an EventBridge rule.
//
// The tag will be propagated to ECS by EventBridge when starting an ECS task based on a matched event.
//
// > Currently, tags are only available when using ECS with EventBridge .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagProperty := &TagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-tag.html
//
type CfnRule_TagProperty struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-tag.html#cfn-events-rule-tag-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the specified tag key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-tag.html#cfn-events-rule-tag-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

