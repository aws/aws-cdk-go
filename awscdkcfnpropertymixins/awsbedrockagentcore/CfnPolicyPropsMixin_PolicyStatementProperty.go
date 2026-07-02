package awsbedrockagentcore


// A policy statement within the AgentCore Policy system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   policyStatementProperty := &PolicyStatementProperty{
//   	Statement: jsii.String("statement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-policystatement.html
//
type CfnPolicyPropsMixin_PolicyStatementProperty struct {
	// The policy statement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-policystatement.html#cfn-bedrockagentcore-policy-policystatement-statement
	//
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}

