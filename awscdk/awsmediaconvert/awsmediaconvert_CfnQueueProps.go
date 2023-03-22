package awsmediaconvert


// Properties for defining a `CfnQueue`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var tags interface{}
//
//   cfnQueueProps := &cfnQueueProps{
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	pricingPlan: jsii.String("pricingPlan"),
//   	status: jsii.String("status"),
//   	tags: tags,
//   }
//
type CfnQueueProps struct {
	// Optional.
	//
	// A description of the queue that you are creating.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the queue that you are creating.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When you use AWS CloudFormation , you can create only on-demand queues.
	//
	// Therefore, always set `PricingPlan` to the value "ON_DEMAND" when declaring an AWS::MediaConvert::Queue in your AWS CloudFormation template.
	//
	// To create a reserved queue, use the AWS Elemental MediaConvert console at https://console.aws.amazon.com/mediaconvert to set up a contract. For more information, see [Working with AWS Elemental MediaConvert Queues](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html) in the ** .
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// Initial state of the queue.
	//
	// Queues can be either ACTIVE or PAUSED. If you create a paused queue, then jobs that you send to that queue won't begin.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

