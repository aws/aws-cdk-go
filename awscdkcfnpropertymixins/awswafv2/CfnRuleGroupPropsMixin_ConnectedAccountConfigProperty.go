package awswafv2


// Configuration for a single connected account and the chains it accepts payment on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectedAccountConfigProperty := &ConnectedAccountConfigProperty{
//   	ConnectedAccountArn: jsii.String("connectedAccountArn"),
//   	PaymentNetworks: []interface{}{
//   		&ConnectedAccountChainConfigProperty{
//   			Chain: jsii.String("chain"),
//   			Prices: []interface{}{
//   				&ConnectedAccountPriceProperty{
//   					Amount: jsii.String("amount"),
//   					Currency: jsii.String("currency"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountconfig.html
//
type CfnRuleGroupPropsMixin_ConnectedAccountConfigProperty struct {
	// The ARN of the connected account that provisions deposit addresses for this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountconfig.html#cfn-wafv2-rulegroup-connectedaccountconfig-connectedaccountarn
	//
	ConnectedAccountArn *string `field:"optional" json:"connectedAccountArn" yaml:"connectedAccountArn"`
	// List of per-chain configurations for a connected account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountconfig.html#cfn-wafv2-rulegroup-connectedaccountconfig-paymentnetworks
	//
	PaymentNetworks interface{} `field:"optional" json:"paymentNetworks" yaml:"paymentNetworks"`
}

