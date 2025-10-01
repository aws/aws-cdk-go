package awsgamelift


// A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an instance in a fleet.
//
// New game sessions are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges. Fleets with custom game builds must have permissions explicitly set. For Realtime Servers fleets, GameLift automatically opens two port ranges, one for TCP messaging and one for UDP.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ippermission.html
//
type CfnFleet_IpPermissionProperty struct {
	// A starting value for a range of allowed port numbers.
	//
	// For fleets using Linux builds, only ports `22` and `1026-60000` are valid.
	//
	// For fleets using Windows builds, only ports `1026-60000` are valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ippermission.html#cfn-gamelift-fleet-ippermission-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: " `000.000.000.000/[subnet mask]` " or optionally the shortened version " `0.0.0.0/[subnet mask]` ".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ippermission.html#cfn-gamelift-fleet-ippermission-iprange
	//
	IpRange *string `field:"required" json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ippermission.html#cfn-gamelift-fleet-ippermission-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than `FromPort` .
	//
	// For fleets using Linux builds, only ports `22` and `1026-60000` are valid.
	//
	// For fleets using Windows builds, only ports `1026-60000` are valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-ippermission.html#cfn-gamelift-fleet-ippermission-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

