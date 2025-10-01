package awscontroltower


// A reference to a EnabledControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   enabledControlReference := &EnabledControlReference{
//   	ControlIdentifier: jsii.String("controlIdentifier"),
//   	TargetIdentifier: jsii.String("targetIdentifier"),
//   }
//
type EnabledControlReference struct {
	// The ControlIdentifier of the EnabledControl resource.
	ControlIdentifier *string `field:"required" json:"controlIdentifier" yaml:"controlIdentifier"`
	// The TargetIdentifier of the EnabledControl resource.
	TargetIdentifier *string `field:"required" json:"targetIdentifier" yaml:"targetIdentifier"`
}

