package awsdeadline


// The range of memory in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryMiBRangeProperty := &MemoryMiBRangeProperty{
//   	Min: jsii.Number(123),
//
//   	// the properties below are optional
//   	Max: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-memorymibrange.html
//
type CfnFleet_MemoryMiBRangeProperty struct {
	// The minimum amount of memory (in MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-memorymibrange.html#cfn-deadline-fleet-memorymibrange-min
	//
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// The maximum amount of memory (in MiB).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-fleet-memorymibrange.html#cfn-deadline-fleet-memorymibrange-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
}

