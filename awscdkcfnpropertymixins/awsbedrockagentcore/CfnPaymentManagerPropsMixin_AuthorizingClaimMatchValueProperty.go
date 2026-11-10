package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizingClaimMatchValueProperty := &AuthorizingClaimMatchValueProperty{
//   	ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   	ClaimMatchValue: &ClaimMatchValueProperty{
//   		MatchValueString: jsii.String("matchValueString"),
//   		MatchValueStringList: []*string{
//   			jsii.String("matchValueStringList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue.html
//
type CfnPaymentManagerPropsMixin_AuthorizingClaimMatchValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"optional" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"optional" json:"claimMatchValue" yaml:"claimMatchValue"`
}

