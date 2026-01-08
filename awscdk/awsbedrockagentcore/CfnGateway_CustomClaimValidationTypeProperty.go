package awsbedrockagentcore


// Required custom claim.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customClaimValidationTypeProperty := &CustomClaimValidationTypeProperty{
//   	AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   		ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   		ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   			MatchValueString: jsii.String("matchValueString"),
//   			MatchValueStringList: []*string{
//   				jsii.String("matchValueStringList"),
//   			},
//   		},
//   	},
//   	InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   	InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customclaimvalidationtype.html
//
type CfnGateway_CustomClaimValidationTypeProperty struct {
	// The value or values in the custom claim to match and relationship of match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customclaimvalidationtype.html#cfn-bedrockagentcore-gateway-customclaimvalidationtype-authorizingclaimmatchvalue
	//
	AuthorizingClaimMatchValue interface{} `field:"required" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
	// The name of the custom claim to validate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customclaimvalidationtype.html#cfn-bedrockagentcore-gateway-customclaimvalidationtype-inboundtokenclaimname
	//
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// Token claim data type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-customclaimvalidationtype.html#cfn-bedrockagentcore-gateway-customclaimvalidationtype-inboundtokenclaimvaluetype
	//
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
}

