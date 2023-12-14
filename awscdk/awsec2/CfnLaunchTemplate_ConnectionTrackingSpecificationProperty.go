package awsec2


// A security group connection tracking specification that enables you to set the idle timeout for connection tracking on an Elastic network interface.
//
// For more information, see [Connection tracking timeouts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-connection-tracking.html#connection-tracking-timeouts) in the *Amazon Elastic Compute Cloud User Guide* .
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
	// Timeout (in seconds) for idle TCP connections in an established state.
	//
	// Min: 60 seconds. Max: 432000 seconds (5 days). Default: 432000 seconds. Recommended: Less than 432000 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-tcpestablishedtimeout
	//
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Timeout (in seconds) for idle UDP flows classified as streams which have seen more than one request-response transaction.
	//
	// Min: 60 seconds. Max: 180 seconds (3 minutes). Default: 180 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-udpstreamtimeout
	//
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Timeout (in seconds) for idle UDP flows that have seen traffic only in a single direction or a single request-response transaction.
	//
	// Min: 30 seconds. Max: 60 seconds. Default: 30 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-connectiontrackingspecification.html#cfn-ec2-launchtemplate-connectiontrackingspecification-udptimeout
	//
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

