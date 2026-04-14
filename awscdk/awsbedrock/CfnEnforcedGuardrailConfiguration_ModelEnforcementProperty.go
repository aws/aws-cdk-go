package awsbedrock


// Model-specific information for the enforced guardrail configuration.
//
// If not present, the configuration is enforced on all models.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modelEnforcementProperty := &ModelEnforcementProperty{
//   	ExcludedModels: []*string{
//   		jsii.String("excludedModels"),
//   	},
//   	IncludedModels: []*string{
//   		jsii.String("includedModels"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-modelenforcement.html
//
type CfnEnforcedGuardrailConfiguration_ModelEnforcementProperty struct {
	// Models to exclude from enforcement.
	//
	// If a model is in both lists, it is excluded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-modelenforcement.html#cfn-bedrock-enforcedguardrailconfiguration-modelenforcement-excludedmodels
	//
	ExcludedModels *[]*string `field:"required" json:"excludedModels" yaml:"excludedModels"`
	// Models to enforce the guardrail on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-enforcedguardrailconfiguration-modelenforcement.html#cfn-bedrock-enforcedguardrailconfiguration-modelenforcement-includedmodels
	//
	IncludedModels *[]*string `field:"required" json:"includedModels" yaml:"includedModels"`
}

