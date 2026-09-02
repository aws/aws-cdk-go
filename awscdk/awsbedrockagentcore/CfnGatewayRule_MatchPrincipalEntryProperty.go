package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchPrincipalEntryProperty := &MatchPrincipalEntryProperty{
//   	IamPrincipal: &IamPrincipalProperty{
//   		Arn: jsii.String("arn"),
//
//   		// the properties below are optional
//   		Operator: jsii.String("operator"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchprincipalentry.html
//
type CfnGatewayRule_MatchPrincipalEntryProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-matchprincipalentry.html#cfn-bedrockagentcore-gatewayrule-matchprincipalentry-iamprincipal
	//
	IamPrincipal interface{} `field:"required" json:"iamPrincipal" yaml:"iamPrincipal"`
}

