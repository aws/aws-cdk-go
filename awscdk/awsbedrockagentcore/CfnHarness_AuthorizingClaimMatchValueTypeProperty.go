package awsbedrockagentcore


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizingclaimmatchvaluetype.html
//
type CfnHarness_AuthorizingClaimMatchValueTypeProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-harness-authorizingclaimmatchvaluetype-claimmatchoperator
	//
	ClaimMatchOperator *string `field:"required" json:"claimMatchOperator" yaml:"claimMatchOperator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-authorizingclaimmatchvaluetype.html#cfn-bedrockagentcore-harness-authorizingclaimmatchvaluetype-claimmatchvalue
	//
	ClaimMatchValue interface{} `field:"required" json:"claimMatchValue" yaml:"claimMatchValue"`
}

