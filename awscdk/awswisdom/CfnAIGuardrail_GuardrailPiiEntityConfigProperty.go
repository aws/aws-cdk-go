package awswisdom


// Pii entity configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailPiiEntityConfigProperty := &GuardrailPiiEntityConfigProperty{
//   	Action: jsii.String("action"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailpiientityconfig.html
//
type CfnAIGuardrail_GuardrailPiiEntityConfigProperty struct {
	// Options for sensitive information action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailpiientityconfig.html#cfn-wisdom-aiguardrail-guardrailpiientityconfig-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The currently supported PII entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiguardrail-guardrailpiientityconfig.html#cfn-wisdom-aiguardrail-guardrailpiientityconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

