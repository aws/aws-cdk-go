package awsbedrockagentcore


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html
//
type CfnHarness_CustomClaimValidationTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-authorizingclaimmatchvalue
	//
	AuthorizingClaimMatchValue interface{} `field:"required" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-inboundtokenclaimname
	//
	InboundTokenClaimName *string `field:"required" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-inboundtokenclaimvaluetype
	//
	InboundTokenClaimValueType *string `field:"required" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
}

