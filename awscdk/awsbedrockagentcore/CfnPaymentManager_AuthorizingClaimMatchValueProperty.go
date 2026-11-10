package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnPaymentManager_AuthorizingClaimMatchValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"required" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvalue-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"required" json:"claimMatchValue" yaml:"claimMatchValue"`
}

