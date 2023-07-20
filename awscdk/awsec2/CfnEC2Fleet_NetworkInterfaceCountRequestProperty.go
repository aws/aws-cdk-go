package awsec2


// The minimum and maximum number of network interfaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceCountRequestProperty := &NetworkInterfaceCountRequestProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacecountrequest.html
//
type CfnEC2Fleet_NetworkInterfaceCountRequestProperty struct {
	// The maximum number of network interfaces.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacecountrequest.html#cfn-ec2-ec2fleet-networkinterfacecountrequest-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum number of network interfaces.
	//
	// To specify no minimum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-networkinterfacecountrequest.html#cfn-ec2-ec2fleet-networkinterfacecountrequest-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

