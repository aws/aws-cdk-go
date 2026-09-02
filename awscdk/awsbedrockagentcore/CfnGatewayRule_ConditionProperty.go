package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &ConditionProperty{
//   	MatchPaths: &MatchPathsProperty{
//   		AnyOf: []*string{
//   			jsii.String("anyOf"),
//   		},
//   	},
//   	MatchPrincipals: &MatchPrincipalsProperty{
//   		AnyOf: []interface{}{
//   			&MatchPrincipalEntryProperty{
//   				IamPrincipal: &IamPrincipalProperty{
//   					Arn: jsii.String("arn"),
//
//   					// the properties below are optional
//   					Operator: jsii.String("operator"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-condition.html
//
type CfnGatewayRule_ConditionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-condition.html#cfn-bedrockagentcore-gatewayrule-condition-matchpaths
	//
	MatchPaths interface{} `field:"optional" json:"matchPaths" yaml:"matchPaths"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-condition.html#cfn-bedrockagentcore-gatewayrule-condition-matchprincipals
	//
	MatchPrincipals interface{} `field:"optional" json:"matchPrincipals" yaml:"matchPrincipals"`
}

