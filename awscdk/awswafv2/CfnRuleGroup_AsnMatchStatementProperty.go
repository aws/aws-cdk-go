package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   asnMatchStatementProperty := &AsnMatchStatementProperty{
//   	AsnList: []interface{}{
//   		jsii.Number(123),
//   	},
//   	ForwardedIpConfig: &ForwardedIPConfigurationProperty{
//   		FallbackBehavior: jsii.String("fallbackBehavior"),
//   		HeaderName: jsii.String("headerName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html
//
type CfnRuleGroup_AsnMatchStatementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html#cfn-wafv2-rulegroup-asnmatchstatement-asnlist
	//
	AsnList interface{} `field:"optional" json:"asnList" yaml:"asnList"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-asnmatchstatement.html#cfn-wafv2-rulegroup-asnmatchstatement-forwardedipconfig
	//
	ForwardedIpConfig interface{} `field:"optional" json:"forwardedIpConfig" yaml:"forwardedIpConfig"`
}

