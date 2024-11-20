package awsgamelift


// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift hosting resource.
//
// New game sessions that are started on the fleet are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges. For fleets created with a custom game server, the ranges reflect the server's game session assignments. For Realtime Servers fleets, Amazon GameLift automatically opens two port ranges, one for TCP messaging and one for UDP, for use by the Realtime servers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipPermissionProperty := &IpPermissionProperty{
//   	FromPort: jsii.Number(123),
//   	IpRange: jsii.String("ipRange"),
//   	Protocol: jsii.String("protocol"),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html
//
type CfnContainerFleet_IpPermissionProperty struct {
	// A starting value for a range of allowed port numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: "000.000.000.000/[subnet mask]" or optionally the shortened version "0.0.0.0/[subnet mask]".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-iprange
	//
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be higher than FromPort.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

