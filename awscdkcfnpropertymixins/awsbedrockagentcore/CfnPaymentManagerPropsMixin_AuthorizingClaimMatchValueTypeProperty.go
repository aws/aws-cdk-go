package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   authorizingClaimMatchValueTypeProperty := &AuthorizingClaimMatchValueTypeProperty{
//   	ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   	ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   		MatchValueString: jsii.String("matchValueString"),
//   		MatchValueStringList: []*string{
//   			jsii.String("matchValueStringList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvaluetype.html
//
type CfnPaymentManagerPropsMixin_AuthorizingClaimMatchValueTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvaluetype-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"optional" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-paymentmanager-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-paymentmanager-authorizingclaimmatchvaluetype-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"optional" json:"claimMatchValue" yaml:"claimMatchValue"`
}

