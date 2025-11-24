package mixinsawswisdom


// Sensitive information policy configuration for a guardrail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aIGuardrailSensitiveInformationPolicyConfigProperty := &AIGuardrailSensitiveInformationPolicyConfigProperty{
//   	PiiEntitiesConfig: []interface{}{
//   		&GuardrailPiiEntityConfigProperty{
//   			Action: jsii.String("action"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	RegexesConfig: []interface{}{
//   		&GuardrailRegexConfigProperty{
//   			Action: jsii.String("action"),
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			Pattern: jsii.String("pattern"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig.html
//
type CfnAIGuardrailPropsMixin_AIGuardrailSensitiveInformationPolicyConfigProperty struct {
	// List of entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig-piientitiesconfig
	//
	PiiEntitiesConfig interface{} `field:"optional" json:"piiEntitiesConfig" yaml:"piiEntitiesConfig"`
	// List of regex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig.html#cfn-wisdom-aiguardrail-aiguardrailsensitiveinformationpolicyconfig-regexesconfig
	//
	RegexesConfig interface{} `field:"optional" json:"regexesConfig" yaml:"regexesConfig"`
}

