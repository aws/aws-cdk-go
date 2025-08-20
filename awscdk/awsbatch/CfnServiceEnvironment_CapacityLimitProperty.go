package awsbatch


// Defines the capacity limit for a service environment.
//
// This structure specifies the maximum amount of resources that can be used by service jobs in the environment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capacityLimitProperty := &CapacityLimitProperty{
//   	CapacityUnit: jsii.String("capacityUnit"),
//   	MaxCapacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-serviceenvironment-capacitylimit.html
//
type CfnServiceEnvironment_CapacityLimitProperty struct {
	// The unit of measure for the capacity limit.
	//
	// This defines how the maxCapacity value should be interpreted. For `SAGEMAKER_TRAINING` jobs, use `NUM_INSTANCES` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-serviceenvironment-capacitylimit.html#cfn-batch-serviceenvironment-capacitylimit-capacityunit
	//
	CapacityUnit *string `field:"optional" json:"capacityUnit" yaml:"capacityUnit"`
	// The maximum capacity available for the service environment.
	//
	// This value represents the maximum amount resources that can be allocated to service jobs.
	//
	// For example, `maxCapacity=50` , `capacityUnit=NUM_INSTANCES` . This indicates that the maximum number of instances that can be run on this service environment is 50. You could then run 5 SageMaker Training jobs that each use 10 instances. However, if you submit another job that requires 10 instances, it will wait in the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-serviceenvironment-capacitylimit.html#cfn-batch-serviceenvironment-capacitylimit-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
}

