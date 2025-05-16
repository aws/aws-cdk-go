package awsnetworkfirewall


// TCP flags and masks to inspect packets for. This is used in the `RuleGroup.MatchAttributes` specification.
//
// For example:
//
// `"TCPFlags": [ { "Flags": [ "ECE", "SYN" ], "Masks": [ "SYN", "ECE" ] } ]`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tCPFlagFieldProperty := &TCPFlagFieldProperty{
//   	Flags: []*string{
//   		jsii.String("flags"),
//   	},
//
//   	// the properties below are optional
//   	Masks: []*string{
//   		jsii.String("masks"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-tcpflagfield.html
//
type CfnRuleGroup_TCPFlagFieldProperty struct {
	// Used in conjunction with the `Masks` setting to define the flags that must be set and flags that must not be set in order for the packet to match.
	//
	// This setting can only specify values that are also specified in the `Masks` setting.
	//
	// For the flags that are specified in the masks setting, the following must be true for the packet to match:
	//
	// - The ones that are set in this flags setting must be set in the packet.
	// - The ones that are not set in this flags setting must also not be set in the packet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-tcpflagfield.html#cfn-networkfirewall-rulegroup-tcpflagfield-flags
	//
	Flags *[]*string `field:"required" json:"flags" yaml:"flags"`
	// The set of flags to consider in the inspection.
	//
	// To inspect all flags in the valid values list, leave this with no setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-tcpflagfield.html#cfn-networkfirewall-rulegroup-tcpflagfield-masks
	//
	Masks *[]*string `field:"optional" json:"masks" yaml:"masks"`
}

