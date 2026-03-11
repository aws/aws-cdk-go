package awsbedrockagentcore


// Properties for CfnPolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPolicyMixinProps := &CfnPolicyMixinProps{
//   	Definition: &PolicyDefinitionProperty{
//   		Cedar: &CedarPolicyProperty{
//   			Statement: jsii.String("statement"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	PolicyEngineId: jsii.String("policyEngineId"),
//   	ValidationMode: jsii.String("validationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html
//
type CfnPolicyMixinProps struct {
	// The definition structure for policies.
	//
	// Encapsulates different policy formats.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-definition
	//
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// A human-readable description of the policy's purpose and functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The customer-assigned immutable name for the policy.
	//
	// Must be unique within the policy engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The identifier of the policy engine which contains this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-policyengineid
	//
	PolicyEngineId *string `field:"optional" json:"policyEngineId" yaml:"policyEngineId"`
	// The validation mode for the policy.
	//
	// Determines how Cedar analyzer validation results are handled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-validationmode
	//
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

