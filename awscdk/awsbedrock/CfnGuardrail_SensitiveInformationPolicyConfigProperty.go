package awsbedrock


// Contains details about PII entities and regular expressions to configure for the guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sensitiveInformationPolicyConfigProperty := &SensitiveInformationPolicyConfigProperty{
//   	PiiEntitiesConfig: []interface{}{
//   		&PiiEntityConfigProperty{
//   			Action: jsii.String("action"),
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	RegexesConfig: []interface{}{
//   		&RegexConfigProperty{
//   			Action: jsii.String("action"),
//   			Name: jsii.String("name"),
//   			Pattern: jsii.String("pattern"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			InputAction: jsii.String("inputAction"),
//   			InputEnabled: jsii.Boolean(false),
//   			OutputAction: jsii.String("outputAction"),
//   			OutputEnabled: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html
//
type CfnGuardrail_SensitiveInformationPolicyConfigProperty struct {
	// A list of PII entities to configure to the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html#cfn-bedrock-guardrail-sensitiveinformationpolicyconfig-piientitiesconfig
	//
	PiiEntitiesConfig interface{} `field:"optional" json:"piiEntitiesConfig" yaml:"piiEntitiesConfig"`
	// A list of regular expressions to configure to the guardrail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-sensitiveinformationpolicyconfig.html#cfn-bedrock-guardrail-sensitiveinformationpolicyconfig-regexesconfig
	//
	RegexesConfig interface{} `field:"optional" json:"regexesConfig" yaml:"regexesConfig"`
}

