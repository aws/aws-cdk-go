package interfacesawsbedrock


// A reference to a EnforcedGuardrailConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enforcedGuardrailConfigurationReference := &EnforcedGuardrailConfigurationReference{
//   	ConfigId: jsii.String("configId"),
//   }
//
type EnforcedGuardrailConfigurationReference struct {
	// The ConfigId of the EnforcedGuardrailConfiguration resource.
	ConfigId *string `field:"required" json:"configId" yaml:"configId"`
}

