package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
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
type CfnHarnessPropsMixin_CustomClaimValidationTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-authorizingclaimmatchvalue
	//
	AuthorizingClaimMatchValue interface{} `field:"optional" json:"authorizingClaimMatchValue" yaml:"authorizingClaimMatchValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-inboundtokenclaimname
	//
	InboundTokenClaimName *string `field:"optional" json:"inboundTokenClaimName" yaml:"inboundTokenClaimName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-customclaimvalidationtype.html#cfn-bedrockagentcore-harness-customclaimvalidationtype-inboundtokenclaimvaluetype
	//
	InboundTokenClaimValueType *string `field:"optional" json:"inboundTokenClaimValueType" yaml:"inboundTokenClaimValueType"`
}

