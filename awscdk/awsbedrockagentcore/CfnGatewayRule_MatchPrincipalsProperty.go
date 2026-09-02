package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchPrincipalsProperty := &MatchPrincipalsProperty{
//   	AnyOf: []interface{}{
//   		&MatchPrincipalEntryProperty{
//   			IamPrincipal: &IamPrincipalProperty{
//   				Arn: jsii.String("arn"),
//
//   				// the properties below are optional
//   				Operator: jsii.String("operator"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchprincipals.html
//
type CfnGatewayRule_MatchPrincipalsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchprincipals.html#cfn-bedrockagentcore-gatewayrule-matchprincipals-anyof
	//
	AnyOf interface{} `field:"required" json:"anyOf" yaml:"anyOf"`
}

