package awsbedrockagentcore


// Properties for defining a `CfnPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPolicyProps := &CfnPolicyProps{
//   	Definition: &PolicyDefinitionProperty{
//   		Cedar: &CedarPolicyProperty{
//   			Statement: jsii.String("statement"),
//   		},
//   		Policy: &PolicyStatementProperty{
//   			Statement: jsii.String("statement"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	PolicyEngineId: jsii.String("policyEngineId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnforcementMode: jsii.String("enforcementMode"),
//   	ValidationMode: jsii.String("validationMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html
//
type CfnPolicyProps struct {
	// The definition structure for policies.
	//
	// Encapsulates different policy formats.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-definition
	//
	Definition interface{} `field:"required" json:"definition" yaml:"definition"`
	// The customer-assigned immutable name for the policy.
	//
	// Must be unique within the policy engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The identifier of the policy engine which contains this policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-policyengineid
	//
	PolicyEngineId *string `field:"required" json:"policyEngineId" yaml:"policyEngineId"`
	// A human-readable description of the policy's purpose and functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the policy contributes to the enforce decision returned to Gateway.
	//
	// LOG_ONLY policies are still evaluated but their decisions are observed only, allowing customers to validate a policy against real traffic before promoting it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-enforcementmode
	//
	// Default: - "ACTIVE".
	//
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
	// The validation mode for the policy.
	//
	// Determines how Cedar analyzer validation results are handled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-policy.html#cfn-bedrockagentcore-policy-validationmode
	//
	ValidationMode *string `field:"optional" json:"validationMode" yaml:"validationMode"`
}

