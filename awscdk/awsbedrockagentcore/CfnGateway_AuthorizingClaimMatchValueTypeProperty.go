package awsbedrockagentcore


// The value or values in the custom claim to match and relationship of match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.html
//
type CfnGateway_AuthorizingClaimMatchValueTypeProperty struct {
	// The relationship between the claim field value and the value or values being matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-gateway-authorizingclaimmatchvaluetype-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"required" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// The value or values in the custom claim to match for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-gateway-authorizingclaimmatchvaluetype-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"required" json:"claimMatchValue" yaml:"claimMatchValue"`
}

