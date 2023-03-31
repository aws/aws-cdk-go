package awscognito


// The type of the configuration to override the risk decision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   riskExceptionConfigurationTypeProperty := &riskExceptionConfigurationTypeProperty{
//   	blockedIpRangeList: []*string{
//   		jsii.String("blockedIpRangeList"),
//   	},
//   	skippedIpRangeList: []*string{
//   		jsii.String("skippedIpRangeList"),
//   	},
//   }
//
type CfnUserPoolRiskConfigurationAttachment_RiskExceptionConfigurationTypeProperty struct {
	// Overrides the risk decision to always block the pre-authentication requests.
	//
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	BlockedIpRangeList *[]*string `field:"optional" json:"blockedIpRangeList" yaml:"blockedIpRangeList"`
	// Risk detection isn't performed on the IP addresses in this range list.
	//
	// The IP range is in CIDR notation.
	SkippedIpRangeList *[]*string `field:"optional" json:"skippedIpRangeList" yaml:"skippedIpRangeList"`
}

