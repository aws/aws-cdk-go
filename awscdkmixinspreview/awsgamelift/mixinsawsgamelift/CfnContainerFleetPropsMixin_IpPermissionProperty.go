package mixinsawsgamelift


// A range of IP addresses and port settings that allow inbound traffic to connect to processes on an instance in a fleet.
//
// Processes are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges.
//
// For Amazon GameLift Servers Realtime fleets, Amazon GameLift Servers automatically opens two port ranges, one for TCP messaging and one for UDP.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnContainerFleetPropsMixin_IpPermissionProperty struct {
	// A starting value for a range of allowed port numbers.
	//
	// For fleets using Linux builds, only ports `22` and `1026-60000` are valid.
	//
	// For fleets using Windows builds, only ports `1026-60000` are valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// A range of allowed IP addresses.
	//
	// This value must be expressed in CIDR notation. Example: " `000.000.000.000/[subnet mask]` " or optionally the shortened version " `0.0.0.0/[subnet mask]` ".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-iprange
	//
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The network communication protocol used by the fleet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// An ending value for a range of allowed port numbers.
	//
	// Port numbers are end-inclusive. This value must be equal to or greater than `FromPort` .
	//
	// For fleets using Linux builds, only ports `22` and `1026-60000` are valid.
	//
	// For fleets using Windows builds, only ports `1026-60000` are valid.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-containerfleet-ippermission.html#cfn-gamelift-containerfleet-ippermission-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

