package awsec2


// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).
//
// > Setting the minimum bandwidth does not guarantee that your instance will achieve the minimum bandwidth. Amazon EC2 will identify instance types that support the specified minimum bandwidth, but the actual bandwidth of your instance might go below the specified minimum at times. For more information, see [Available instance bandwidth](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-network-bandwidth.html#available-instance-bandwidth) in the *Amazon EC2 User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkBandwidthGbpsProperty := &NetworkBandwidthGbpsProperty{
//   	Max: jsii.Number(123),
//   	Min: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkbandwidthgbps.html
//
type CfnLaunchTemplate_NetworkBandwidthGbpsProperty struct {
	// The maximum amount of network bandwidth, in Gbps.
	//
	// To specify no maximum limit, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkbandwidthgbps.html#cfn-ec2-launchtemplate-networkbandwidthgbps-max
	//
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// The minimum amount of network bandwidth, in Gbps.
	//
	// If this parameter is not specified, there is no minimum limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkbandwidthgbps.html#cfn-ec2-launchtemplate-networkbandwidthgbps-min
	//
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

