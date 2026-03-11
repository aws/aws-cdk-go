package awsbedrockagentcore


// A Cedar policy statement within the AgentCore Policy system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cedarPolicyProperty := &CedarPolicyProperty{
//   	Statement: jsii.String("statement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-cedarpolicy.html
//
type CfnPolicyPropsMixin_CedarPolicyProperty struct {
	// The Cedar policy statement that defines the authorization logic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-cedarpolicy.html#cfn-bedrockagentcore-policy-cedarpolicy-statement
	//
	Statement *string `field:"optional" json:"statement" yaml:"statement"`
}

