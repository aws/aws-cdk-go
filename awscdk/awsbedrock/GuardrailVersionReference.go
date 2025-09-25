package awsbedrock


// A reference to a GuardrailVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   guardrailVersionReference := &GuardrailVersionReference{
//   	GuardrailId: jsii.String("guardrailId"),
//   	Version: jsii.String("version"),
//   }
//
type GuardrailVersionReference struct {
	// The GuardrailId of the GuardrailVersion resource.
	GuardrailId *string `field:"required" json:"guardrailId" yaml:"guardrailId"`
	// The Version of the GuardrailVersion resource.
	Version *string `field:"required" json:"version" yaml:"version"`
}

