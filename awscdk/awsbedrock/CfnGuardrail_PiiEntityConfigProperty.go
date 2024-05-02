package awsbedrock


// Pii entity configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   piiEntityConfigProperty := &PiiEntityConfigProperty{
//   	Action: jsii.String("action"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html
//
type CfnGuardrail_PiiEntityConfigProperty struct {
	// Options for sensitive information action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html#cfn-bedrock-guardrail-piientityconfig-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// The currently supported PII entities.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-guardrail-piientityconfig.html#cfn-bedrock-guardrail-piientityconfig-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

