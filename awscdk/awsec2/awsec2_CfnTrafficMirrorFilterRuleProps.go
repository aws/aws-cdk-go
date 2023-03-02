package awsec2


// Properties for defining a `CfnTrafficMirrorFilterRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTrafficMirrorFilterRuleProps := &cfnTrafficMirrorFilterRuleProps{
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	ruleAction: jsii.String("ruleAction"),
//   	ruleNumber: jsii.Number(123),
//   	sourceCidrBlock: jsii.String("sourceCidrBlock"),
//   	trafficDirection: jsii.String("trafficDirection"),
//   	trafficMirrorFilterId: jsii.String("trafficMirrorFilterId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	destinationPortRange: &trafficMirrorPortRangeProperty{
//   		fromPort: jsii.Number(123),
//   		toPort: jsii.Number(123),
//   	},
//   	protocol: jsii.Number(123),
//   	sourcePortRange: &trafficMirrorPortRangeProperty{
//   		fromPort: jsii.Number(123),
//   		toPort: jsii.Number(123),
//   	},
//   }
//
type CfnTrafficMirrorFilterRuleProps struct {
	// The destination CIDR block to assign to the Traffic Mirror rule.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The action to take on the filtered traffic.
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// The number of the Traffic Mirror rule.
	//
	// This number must be unique for each Traffic Mirror rule in a given direction. The rules are processed in ascending order by rule number.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// The source CIDR block to assign to the Traffic Mirror rule.
	SourceCidrBlock *string `field:"required" json:"sourceCidrBlock" yaml:"sourceCidrBlock"`
	// The type of traffic.
	TrafficDirection *string `field:"required" json:"trafficDirection" yaml:"trafficDirection"`
	// The ID of the filter that this rule is associated with.
	TrafficMirrorFilterId *string `field:"required" json:"trafficMirrorFilterId" yaml:"trafficMirrorFilterId"`
	// The description of the Traffic Mirror rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The destination port range.
	DestinationPortRange interface{} `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The protocol, for example UDP, to assign to the Traffic Mirror rule.
	//
	// For information about the protocol value, see [Protocol Numbers](https://docs.aws.amazon.com/https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) on the Internet Assigned Numbers Authority (IANA) website.
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// The source port range.
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

