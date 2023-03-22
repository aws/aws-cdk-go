package awsmediaconvert


// Optional.
//
// Configuration for a destination queue to which the job can hop once a customer-defined minimum wait time has passed. For more information, see [Setting Up Queue Hopping to Avoid Long Waits](https://docs.aws.amazon.com/mediaconvert/latest/ug/setting-up-queue-hopping-to-avoid-long-waits.html) in the *AWS Elemental MediaConvert User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hopDestinationProperty := &hopDestinationProperty{
//   	priority: jsii.Number(123),
//   	queue: jsii.String("queue"),
//   	waitMinutes: jsii.Number(123),
//   }
//
type CfnJobTemplate_HopDestinationProperty struct {
	// Optional.
	//
	// When you set up a job to use queue hopping, you can specify a different relative priority for the job in the destination queue. If you don't specify, the relative priority will remain the same as in the previous queue.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Optional unless the job is submitted on the default queue.
	//
	// When you set up a job to use queue hopping, you can specify a destination queue. This queue cannot be the original queue to which the job is submitted. If the original queue isn't the default queue and you don't specify the destination queue, the job will move to the default queue.
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
	// Required for setting up a job to use queue hopping.
	//
	// Minimum wait time in minutes until the job can hop to the destination queue. Valid range is 1 to 1440 minutes, inclusive.
	WaitMinutes *float64 `field:"optional" json:"waitMinutes" yaml:"waitMinutes"`
}

