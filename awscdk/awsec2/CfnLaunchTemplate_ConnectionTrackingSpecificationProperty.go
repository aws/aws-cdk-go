package awsec2


// Allows customer to specify Connection Tracking options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionTrackingSpecificationProperty := &ConnectionTrackingSpecificationProperty{
//   	TcpEstablishedTimeout: jsii.Number(123),
//   	UdpStreamTimeout: jsii.Number(123),
//   	UdpTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html
//
type CfnLaunchTemplate_ConnectionTrackingSpecificationProperty struct {
	// Integer value for TCP Established Timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-tcpestablishedtimeout
	//
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Integer value for UDP Stream Timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-udpstreamtimeout
	//
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Integer value for UDP Timeout.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-udptimeout
	//
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

