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
//   cfnQueueProps := &CfnQueueProps{
//   	ConcurrentJobs: jsii.Number(123),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PricingPlan: jsii.String("pricingPlan"),
//   	Status: jsii.String("status"),
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html
//
type CfnQueueProps struct {
	// Specify the maximum number of jobs your queue can process concurrently.
	//
	// For on-demand queues, the value you enter is constrained by your service quotas for Maximum concurrent jobs, per on-demand queue and Maximum concurrent jobs, per account. For reserved queues, specify the number of jobs you can process concurrently in your reservation plan instead.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-concurrentjobs
	//
	ConcurrentJobs *float64 `field:"optional" json:"concurrentJobs" yaml:"concurrentJobs"`
	// Optional.
	//
	// A description of the queue that you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the queue that you are creating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When you use CloudFormation , you can create only on-demand queues.
	//
	// Therefore, always set `PricingPlan` to the value "ON_DEMAND" when declaring an AWS::MediaConvert::Queue in your CloudFormation template.
	//
	// To create a reserved queue, use the AWS Elemental MediaConvert console at https://console.aws.amazon.com/mediaconvert to set up a contract. For more information, see [Working with AWS Elemental MediaConvert Queues](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html) in the ** .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-pricingplan
	//
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// Initial state of the queue.
	//
	// Queues can be either ACTIVE or PAUSED. If you create a paused queue, then jobs that you send to that queue won't begin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-queue.html#cfn-mediaconvert-queue-tags
	//
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

