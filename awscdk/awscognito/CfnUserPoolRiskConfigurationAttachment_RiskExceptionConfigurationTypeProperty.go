package awscognito


// The type of the configuration to override the risk decision.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   riskExceptionConfigurationTypeProperty := &RiskExceptionConfigurationTypeProperty{
//   	BlockedIpRangeList: []*string{
//   		jsii.String("blockedIpRangeList"),
//   	},
//   	SkippedIpRangeList: []*string{
//   		jsii.String("skippedIpRangeList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype.html
//
type CfnUserPoolRiskConfigurationAttachment_RiskExceptionConfigurationTypeProperty struct {
	// Overrides the risk decision to always block the pre-authentication requests.
	//
	// The IP range is in CIDR notation, a compact representation of an IP address and its routing prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype-blockediprangelist
	//
	BlockedIpRangeList *[]*string `field:"optional" json:"blockedIpRangeList" yaml:"blockedIpRangeList"`
	// Risk detection isn't performed on the IP addresses in this range list.
	//
	// The IP range is in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype.html#cfn-cognito-userpoolriskconfigurationattachment-riskexceptionconfigurationtype-skippediprangelist
	//
	SkippedIpRangeList *[]*string `field:"optional" json:"skippedIpRangeList" yaml:"skippedIpRangeList"`
}

