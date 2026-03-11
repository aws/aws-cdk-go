package awsbedrockagentcore


// The definition structure for policies.
//
// Encapsulates different policy formats.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   policyDefinitionProperty := &PolicyDefinitionProperty{
//   	Cedar: &CedarPolicyProperty{
//   		Statement: jsii.String("statement"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-policydefinition.html
//
type CfnPolicyPropsMixin_PolicyDefinitionProperty struct {
	// A Cedar policy statement within the AgentCore Policy system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-policy-policydefinition.html#cfn-bedrockagentcore-policy-policydefinition-cedar
	//
	Cedar interface{} `field:"optional" json:"cedar" yaml:"cedar"`
}

