package awswafv2


// Per-chain price configuration for a connected account.
//
// The deposit address is provisioned by the payments provider and is not specified here.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectedAccountChainConfigProperty := &ConnectedAccountChainConfigProperty{
//   	Chain: jsii.String("chain"),
//   	Prices: []interface{}{
//   		&ConnectedAccountPriceProperty{
//   			Amount: jsii.String("amount"),
//   			Currency: jsii.String("currency"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountchainconfig.html
//
type CfnRuleGroup_ConnectedAccountChainConfigProperty struct {
	// The blockchain chain to use with a connected account.
	//
	// Connected accounts are mainnet only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountchainconfig.html#cfn-wafv2-rulegroup-connectedaccountchainconfig-chain
	//
	Chain *string `field:"required" json:"chain" yaml:"chain"`
	// List of connected-account price configurations.
	//
	// A single price is supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafv2-rulegroup-connectedaccountchainconfig.html#cfn-wafv2-rulegroup-connectedaccountchainconfig-prices
	//
	Prices interface{} `field:"required" json:"prices" yaml:"prices"`
}

