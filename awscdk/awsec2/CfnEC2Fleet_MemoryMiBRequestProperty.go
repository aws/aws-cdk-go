package awsec2


// The minimum and maximum amount of memory, in MiB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryMiBRequestProperty := &MemoryMiBRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-memorymibrequest.html
//
type CfnEC2Fleet_MemoryMiBRequestProperty struct {
	// The maximum amount of memory, in MiB.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-memorymibrequest.html#cfn-ec2-ec2fleet-memorymibrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of memory, in MiB.
	//
	// To specify no minimum limit, specify `0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-memorymibrequest.html#cfn-ec2-ec2fleet-memorymibrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

