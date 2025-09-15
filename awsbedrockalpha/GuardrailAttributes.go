package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// ****************************************************************************                     ATTRS FOR IMPORTED CONSTRUCT ***************************************************************************.
//
// Example:
//   var stack stack
//
//   cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
//   })
//   // Import an existing guardrail by ARN
//   importedGuardrail := bedrock.Guardrail_FromGuardrailAttributes(stack, jsii.String("TestGuardrail"), &GuardrailAttributes{
//   	GuardrailArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:guardrail/oygh3o8g7rtl"),
//   	GuardrailVersion: jsii.String("1"),
//   	 //optional
//   	KmsKey: cmk,
//   })
//
// Experimental.
type GuardrailAttributes struct {
	// The ARN of the guardrail.
	//
	// At least one of guardrailArn or guardrailId must be
	// defined in order to initialize a guardrail ref.
	// Experimental.
	GuardrailArn *string `field:"required" json:"guardrailArn" yaml:"guardrailArn"`
	// The version of the guardrail.
	// Default: "DRAFT".
	//
	// Experimental.
	GuardrailVersion *string `field:"optional" json:"guardrailVersion" yaml:"guardrailVersion"`
	// The KMS key of the guardrail if custom encryption is configured.
	// Default: undefined - Means data is encrypted by default with a AWS-managed key.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

