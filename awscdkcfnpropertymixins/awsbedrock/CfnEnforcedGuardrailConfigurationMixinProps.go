package awsbedrock


// Properties for CfnEnforcedGuardrailConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnEnforcedGuardrailConfigurationMixinProps := &CfnEnforcedGuardrailConfigurationMixinProps{
//   	GuardrailIdentifier: jsii.String("guardrailIdentifier"),
//   	GuardrailVersion: jsii.String("guardrailVersion"),
//   	ModelEnforcement: &ModelEnforcementProperty{
//   		ExcludedModels: []*string{
//   			jsii.String("excludedModels"),
//   		},
//   		IncludedModels: []*string{
//   			jsii.String("includedModels"),
//   		},
//   	},
//   	SelectiveContentGuarding: &SelectiveContentGuardingProperty{
//   		Messages: jsii.String("messages"),
//   		System: jsii.String("system"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html
//
type CfnEnforcedGuardrailConfigurationMixinProps struct {
	// Identifier for the guardrail, could be the ID or the ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html#cfn-bedrock-enforcedguardrailconfiguration-guardrailidentifier
	//
	GuardrailIdentifier *string `field:"optional" json:"guardrailIdentifier" yaml:"guardrailIdentifier"`
	// Numerical guardrail version (not DRAFT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html#cfn-bedrock-enforcedguardrailconfiguration-guardrailversion
	//
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
	// Model-specific information for the enforced guardrail configuration.
	//
	// If not present, the configuration is enforced on all models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html#cfn-bedrock-enforcedguardrailconfiguration-modelenforcement
	//
	ModelEnforcement interface{} `field:"optional" json:"modelEnforcement" yaml:"modelEnforcement"`
	// Selective content guarding controls for enforced guardrails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-enforcedguardrailconfiguration.html#cfn-bedrock-enforcedguardrailconfiguration-selectivecontentguarding
	//
	SelectiveContentGuarding interface{} `field:"optional" json:"selectiveContentGuarding" yaml:"selectiveContentGuarding"`
}

