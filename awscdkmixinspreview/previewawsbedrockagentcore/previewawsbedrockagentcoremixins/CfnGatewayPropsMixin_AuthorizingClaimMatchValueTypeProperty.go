package previewawsbedrockagentcoremixins


// The value or values in the custom claim to match and relationship of match.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnGatewayPropsMixin_AuthorizingClaimMatchValueTypeProperty struct {
	// The relationship between the claim field value and the value or values being matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-gateway-authorizingclaimmatchvaluetype-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"optional" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// The value or values in the custom claim to match for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-gateway-authorizingclaimmatchvaluetype-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"optional" json:"claimMatchValue" yaml:"claimMatchValue"`
}

