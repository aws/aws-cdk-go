package awsnetworkfirewall


// A single port range specification.
//
// This is used for source and destination port ranges in the stateless `RuleGroup.MatchAttributes` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &PortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html
//
type CfnRuleGroup_PortRangeProperty struct {
	// The lower limit of the port range.
	//
	// This must be less than or equal to the `ToPort` specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html#cfn-networkfirewall-rulegroup-portrange-fromport
	//
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// The upper limit of the port range.
	//
	// This must be greater than or equal to the `FromPort` specification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-rulegroup-portrange.html#cfn-networkfirewall-rulegroup-portrange-toport
	//
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
}

