package awsdeadline


// Defines the maximum and minimum amount of memory, in MiB, to use for the accelerator.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorTotalMemoryMiBRangeProperty := &AcceleratorTotalMemoryMiBRangeProperty{
//   	Min: jsii.Number(123),
//
//   	// the properties below are optional
//   	Max: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratortotalmemorymibrange.html
//
type CfnFleet_AcceleratorTotalMemoryMiBRangeProperty struct {
	// The minimum amount of memory to use for the accelerator, measured in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratortotalmemorymibrange.html#cfn-deadline-fleet-acceleratortotalmemorymibrange-min
	//
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// The maximum amount of memory to use for the accelerator, measured in MiB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-acceleratortotalmemorymibrange.html#cfn-deadline-fleet-acceleratortotalmemorymibrange-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

