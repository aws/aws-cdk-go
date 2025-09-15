package awsec2


// The minimum and maximum number of vCPUs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vCpuCountRangeRequestProperty := &VCpuCountRangeRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-vcpucountrangerequest.html
//
type CfnEC2Fleet_VCpuCountRangeRequestProperty struct {
	// The maximum number of vCPUs.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-vcpucountrangerequest.html#cfn-ec2-ec2fleet-vcpucountrangerequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of vCPUs.
	//
	// To specify no minimum limit, specify `0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-vcpucountrangerequest.html#cfn-ec2-ec2fleet-vcpucountrangerequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

