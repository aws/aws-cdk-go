package awsbedrock


// A reference to a Guardrail resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailReference := &GuardrailReference{
//   	GuardrailArn: jsii.String("guardrailArn"),
//   }
//
type GuardrailReference struct {
	// The GuardrailArn of the Guardrail resource.
	GuardrailArn *string `field:"required" json:"guardrailArn" yaml:"guardrailArn"`
}

